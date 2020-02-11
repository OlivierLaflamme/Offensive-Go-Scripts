## Basics    

`go run PROGRAM.go`
`go build PROGRAM.go`


# to reduce the binary size by ~30% by stripping information from the binary     
`go build -ldflags "-w -s" PROGRAM.go`

# to build a binary that can run on different architectures 
To cross compile you'll need to set `constraints` -> which means to pass information to the build command about the opperating system for which you'd like your code to compile to. These  `constraints` include `GOOS` (for opperating systems) and `GOARCH` (for architectures) you can introdduce build constraints in 3 different ways.      
