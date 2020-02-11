## Basics    

`go run PROGRAM.go`
`go build PROGRAM.go`

## Building and cross-compiling
# to reduce the binary size by ~30% by stripping information from the binary     
`go build -ldflags "-w -s" PROGRAM.go`

# to build a binary that can run on different architectures 
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
