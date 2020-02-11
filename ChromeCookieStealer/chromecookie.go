package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/levigross/grequests"
	"github.com/tidwall/gjson"
	"golang.org/x/net/websocket"
)

func runchrome() {
	cmd := exec.Command("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome", "--headless", "--user-data-dir=/Users/ADDYOURUSERNAME/Library/Application Support/Google/Chrome/Default", "https://gmail.com", "--remote-debugging-port=9222")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func checkcookies(socketurl string) {
	cookie := `{"id": 1, "method": "Network.getAllCookies"}`
	origin := "http://localhost/"
	ws, err := websocket.Dial(socketurl, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := ws.Write([]byte(cookie)); err != nil {
		log.Fatal(err)
	}

	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])

}

func main() {
	go runchrome()
	resp, err := grequests.Get("http://localhost:9222/json", nil)
	if err != nil {
		panic(err)
	}
	value := gjson.Get(resp.String(), "#.webSocketDebuggerUrl")
	value.ForEach(func(key, value gjson.Result) bool {
		checkcookies(value.String())
		return true
	})

}
