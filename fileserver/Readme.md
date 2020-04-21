# Go fileserver with uploads 
![Build](https://img.shields.io/badge/build-passing-green) ![CC](https://img.shields.io/badge/license-cc--by--sa--4.0--Licence-blue) ![COMMIT](https://img.shields.io/github/last-commit/OlivierLaflamme/Offensive-Go-Scripts)
------------

This is a fileserver written in golang. As of now, this fileserver will not provide encryption or any other security measures besides  logon authentication (if ran with -pass & -user flags configured) - this is just a simple fileserver. The script itself was (kinda but not really) inspired by [UniIsland/SimpleHTTPServerWithUpload.py](https://gist.github.com/UniIsland/3346170) a well know python SimpleHTTPServer 'config' - however my fileserver actually works out of the box because its not from 2012 using deprecated modules.     

As it stands, this fileserver doesn't support multiple file uploads - you can only upload one file at a time so trying to upload an array of files using `curl -X POST http://IP:80 -F "upload[]=@/Users/test/a.exe" -F "upload[]=@/Users/boschko/b.dll" -H "Content-Type: multipart/form-data"` ... Moreover, this is simply not the intent / use case of this script.


Getting Started
------------
For those of you who are new to go and just want to run this ~ **you only have to run `go get github.com/gin-gonic/gin` once.**    

Running: `go get github.com/gin-gonic/gin && go run filserver.go`   
Building: `go get github.com/gin-gonic/gin && go build -ldflags "-w -s" filserver.go`    

**Or build**     
First run: `go get github.com/gin-gonic/gin`    
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

#### Start the go fileserver with authentication:  
![image](https://user-images.githubusercontent.com/25066959/74349193-df423700-4d81-11ea-9b99-fd1a9cf89b37.png)    

#### Login with credentials 
![image](https://user-images.githubusercontent.com/25066959/74278357-157ea880-4ce7-11ea-819a-3df1ddd970f7.png)     

#### Upload file from local machine
![image](https://user-images.githubusercontent.com/25066959/74277255-33e3a480-4ce5-11ea-8479-36177eca439a.png)    

#### View the file on fileserver
![image](https://user-images.githubusercontent.com/25066959/74349293-0993f480-4d82-11ea-9ad7-fba5db12f04a.png)     


Resources
---------

-  https://github.com/gin-gonic/gin
