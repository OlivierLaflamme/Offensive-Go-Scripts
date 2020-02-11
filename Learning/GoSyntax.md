#### Cheatsheet => https://devhints.io/go  && https://github.com/a8m/golang-cheat-sheet

**Primitive data types:** bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, byte, rune, float32, float64, complex64, and complex128      

You typically declare a variable’s type when you define it. If you don’t, the system will automatically infer the variable’s data type. Consider the following examples    
```go
var x = "hello world" \\ keyword var used to define variable x and assignes to it the value of "hello world"
z := int(42)           \\ uses := opperator to ddefine a new variable z and assign to it an integer value 42
```    
4 ways to create a static type in and value type:    
```go
var a int = 10

var b = 10

var c int 
c=10 

d:= 10 
```
