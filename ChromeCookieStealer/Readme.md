Go version of this interesting project => https://github.com/defaultnamehere/cookie_crimes which steals chrome cookies without root. 

Note:
------------
You'll habe to modify line **13** parameter entitled `ADDYOURUSERNAME` to that of your hostname/username so that the path "exists"    

Building:
------------
Since we used a few libraries you'll have to preform the following
* `go get github.com/levigross/grequests`
* `go get github.com/tidwall/gjson`
* `go get golang.org/x/net/websocket`

To build the application run the following:     
```go
go build -ldflags "-w -s" chromecookie.go
```
or build

*   env GOOS=windows GOARCH=386 go build -v chromecookie.go
*   env GOOS=linux GOARCH=amd64 go build -v chromecookie.go

