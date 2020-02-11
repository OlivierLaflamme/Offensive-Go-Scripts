# Uses Go net/http library to distinguish alive hosts from a give list of hosts/urls 

How to build:
`go build -ldflags "-w -s" main.go`

How to use: 
`cat urls.txt | online`
