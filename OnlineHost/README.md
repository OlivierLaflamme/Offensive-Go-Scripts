# OnlineHosts

#### Uses Go net/http library to distinguish alive hosts from a give list of hosts/urls 

Building:
------------
To build the application run the following:     
```go
go build -ldflags "-w -s" main.go
```
or build

*   env GOOS=windows GOARCH=386 go build -v main.go
*   env GOOS=linux GOARCH=amd64 go build -v main.go

Use:
------------
`cat urls.txt | online`
