#### Cheatsheet => https://devhints.io/go  && https://github.com/a8m/golang-cheat-sheet

**Primitive data types:** bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, byte, rune, float32, float64, complex64, and complex128      

You typically declare a variable’s type when you define it. If you don’t, the system will automatically infer the variable’s data type. Consider the following examples    
```go
var x = "hello world" \\ keyword var used to define variable x and assignes to it the value of "hello world"
z := int(42)           \\ uses := opperator to ddefine a new variable z and assign to it an integer value 42
```    
In the preceding example, you explicitly wrap the 42 value in an intcall to force a type on it. You could omit the int call but would have to accept whatever type the system automatically uses for that value. In some cases, this won’t be the type you intended to use. For instance, perhaps you want 42 to be represented as an unsigned integer, rather than an int type, in which case you’d have to explicitly wrap the value    

4 ways to create a static type in and value type:    
```go
var a int = 10

var b = 10

var c int 
c=10 

d:= 10 
```


**Slice and Maps**    
slices are like arrays that you can dynamically resize and pass to runctions more efficiently.   
maps are associative arraysm unordered lists of key/value pairs that allow you to efficiently and quickly look up values for a unique key.    
there are tons of ways to define, initialize, and work with slices and maps the following demonstrates ways to define both a `slice "s"` and a `map "m"` and add elements to both:
```go 
var s = make([]string, 0)
var m = make(map[string]string)
s = append(s, "some string")
m["some key"] = "some value"
```
code uses two built-in-functions `make()` which initializes each variables and `append()` to add a new item to a slice. the last lines adds the key/value pair of `some key` and `some value` to the map m.

**Pointers, Structs, and Interfaces**    
pointers point to a particulat area in memory and allows you to retrieve the value stored there. the following illustrates this:
```go
var count = int(42)
ptr := &count
fmt.Println(*ptr)
*ptr = 100
fmt.Println(count)
```
defines integer `count` -> creates a pointer by using the `&` opperator (this returns the address of `count`) -> dereference the variable (line3) while making the call to `fmt.Println()` to log the value of `count` to stdout -> use the `*` operator to assign a new value to the memory location pointed by ptr (line4) -> because this is the address of `count` variable, the assignment changes the value of that variable -> which you confirm by printing it to the screen     
























