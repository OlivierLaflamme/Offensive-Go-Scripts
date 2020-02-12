# FileServer which supports uploads 

Usage
------------

Example:    
Running: `go run filserver.go`   
Building: `go build -ldflags "-w -s" filserver.go`    

**Or build**

*   env GOOS=windows GOARCH=386 go build -v filserver.go
*   env GOOS=linux GOARCH=amd64 go build -v filserver.go

**Or Makefile**

Run: `Make` will build for all OS compatibilities 

Usage  
----------

|Parameter     |Description  |
|-----------    |-------------------------------------------------------|
|-h, --help     | show help message and exit                            |
|-dir (string)  | file dir (default "./")                               |
|-pass (string) | password                                              |
|-port (string) | http port (default "80")                              |
|-user (string) | username                                              |


Example
----------

### Start the go fileserver with authentication: 
![image](https://user-images.githubusercontent.com/25066959/74278421-32b37700-4ce7-11ea-9233-009e0ce04d88.png)    

### Login with credentials 
![image](https://user-images.githubusercontent.com/25066959/74278357-157ea880-4ce7-11ea-819a-3df1ddd970f7.png)     

### Upload file from local machine
![image](https://user-images.githubusercontent.com/25066959/74277255-33e3a480-4ce5-11ea-8479-36177eca439a.png)    

### View the file on fileserver
![image](https://user-images.githubusercontent.com/25066959/74278609-7c9c5d00-4ce7-11ea-8e71-99e8971f81ac.png)


Resources
---------

-  https://github.com/gin-gonic/gin
