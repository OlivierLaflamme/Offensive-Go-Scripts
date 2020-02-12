package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)
 
var (
	port     string
	username string
	password string
	dir      string
	head     = `<html>
	<form action="/newdir/" method="POST">
		<h1>Brought to you by: Boschko</h1>
		<input type="text" name="dirname">
		<input type="submit" value="New directory">
	</form>
	<form action="/" enctype="multipart/form-data" method="POST">
		<input type="file" name="files" multiple="multiple" />
		<input type="submit" value="Upload" />
	</form>
	<script>
		window.onload = function() {
			var prelist = document.getElementsByTagName("pre");
			if (prelist.length > 0) {
				var pre = prelist[0];
				var alist = pre.getElementsByTagName("a");
				if (alist.length > 0) {
					for (var i = 0; i < alist.length; i++) {
						var a = alist[i];
						var input = document.createElement("input");
						input.type = "checkbox";
						input.name = "list[]";
						input.value = a.text;
						pre.insertBefore(input, a);
					}
					var input = document.createElement("input");
					input.type = "submit";
					input.value = "delete selected";
					pre.parentNode.append(input);
				}
			}
		}
    </script>
	<form action="/del/" method="POST">
`
	errFmt = `<font color="red">%s</font><br />
	<input type="button" value="return" onclick="history.back()">`
)

type Files struct {
	List []string `form:"list[]"`
}

func init() {
	flag.StringVar(&port, "port", "80", "http port")
	flag.StringVar(&username, "user", "", "username")
	flag.StringVar(&password, "pass", "", "password")
	flag.StringVar(&dir, "dir", "./", "file dir")
	flag.Parse()
}

func uploadFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, fmt.Sprintf(errFmt, "Error: "+err.Error()))
		return
	}

	p := getPath(c.GetHeader("referer"))
	files := form.File["files"]
	for _, file := range files {
		err = c.SaveUploadedFile(file, path.Join(dir, p, path.Clean("/"+file.Filename)))
		if err != nil {
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(200, fmt.Sprintf(errFmt, file.Filename+" upload failed<br />Error: "+err.Error()))
			return
		}
	}
	c.Redirect(302, p)
}

func delFile(c *gin.Context) {
	var files Files
	err := c.ShouldBind(&files)
	if err != nil {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, fmt.Sprintf(errFmt, "Error: "+err.Error()))
		return
	}

	p := getPath(c.GetHeader("referer"))
	for _, file := range files.List {
		err = os.RemoveAll(path.Join(dir, p, path.Clean("/"+file)))
		if err != nil {
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(200, fmt.Sprintf(errFmt, file+" failed to delete<br />Error: "+err.Error()))
			return
		}
	}
	c.Redirect(302, p)
}

func newDir(c *gin.Context) {
	p := getPath(c.GetHeader("referer"))
	dirname := c.PostForm("dirname")
	if dirname != "" {
		err := os.MkdirAll(path.Join(dir, p, path.Clean("/"+dirname)), os.ModeDir)
		if err != nil {
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(200, fmt.Sprintf(errFmt, "Error: "+err.Error()))
			return
		}
	}
	c.Redirect(302, p)
}

func getPath(referer string) string {
	var p = "/"
	u, err := url.ParseRequestURI(referer)
	if err == nil {
		// remove ../
		p = path.Clean("/" + u.Path)
		if p != "/" {
			p += "/"
		}
	}
	return p
}

func writeHead(c *gin.Context) {
	if strings.HasSuffix(c.Request.URL.Path, "/") {
		c.Writer.WriteString(head)
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default() 

	var authGroup *gin.RouterGroup
	if username != "" {
		authGroup = r.Group("/", gin.BasicAuth(gin.Accounts{
			username: password,
		}))
	} else {
		authGroup = r.Group("/")
	}

	getGroup := authGroup.Group("/", writeHead)

	getGroup.StaticFS("/", gin.Dir(dir, true))
	authGroup.POST("/", uploadFile)
	authGroup.POST("/del/", delFile)
	authGroup.POST("/newdir/", newDir)

	err := r.Run(":" + port)
	if err != nil {
		fmt.Println(err)
	}
}
