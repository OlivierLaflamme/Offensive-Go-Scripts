### Go Basics    

`go run PROGRAM.go`   
`go build PROGRAM.go`

===============================================================================


#### [BUILD] and cross-compiling
##### to reduce the binary size by ~30% by stripping information from the binary     
`go build -ldflags "-w -s" PROGRAM.go`

##### to build a binary that can run on different architectures 
To cross compile you'll need to set `constraints` -> which means to pass information to the build command about the opperating system for which you'd like your code to compile to. These  `constraints` include `GOOS` (for opperating systems) and `GOARCH` (for architectures) you can introdduce build constraints in 3 different ways.      
1. via the commandline        
2. code comments     
3. a file suffix naming convention    

Example cross compile so that PROGRAM.go runs on a Linux 64-bit architecture, well accomplish this via the `GOOS` and `GOARCH` constrainsts when running the build command: 
```go
$-> GOOS="Linux" GOARCH="amd64" go build PROGRAM.go
$-> file PROGRAM
$<- PROGRAM: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, not stripped
```  
===============================================================================


#### [HELP] For information/documentation about a package function, methode, commands use `go doc`
`$-> go doc fmt.Println`   
`$<- func Println(a ...interface{}) (n int, err error) Println formats using the default form....`    

===============================================================================


#### [GET PACKAGES] Some programs require third-party packages, to obtaine package sourcecodes use `go get`
Example:  we import the stacktitan/ldapauth package
```go 
package main
import (
"fmt"
"net/http"
"github.com/stacktitan/ldapauth"
)
```    
even though you've imported the package you can access it untill you run `go get github.com/stacktitan/ldapauth` which will download the actual package and place it within the $GOPATH/src directory verify (on linux) with `tree src/github.com/stacktitan/`    
##### NOTE
problems can arrise if dependancy packages receive updates that break backwardds compatibility; there are two seperate tools `dep` and `dmod` to lock dependancies in order to prevent backwards compatibility issues.    

===============================================================================

`fmt` command automatically formats your source code -> `go fmt /path/to/your/package` will stype your code enforcing the proper use of line breakes, indents...     

`golint` reports style mistakes such as missing comments, variable naming that doesnt follow convention....     

`https://play.golang.org/` provides web based front end 
















