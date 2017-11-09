

## Why Go?



## Reference

[The docs[(https://golang.org/doc/)
[The spec](https://golang.org/ref/spec)


### Type System

[Understanding the Go Type System](https://thenewstack.io/understanding-golang-type-system/)


### Get Up and Running

[Go command](https://golang.org/cmd/go/)


### Error Handling

[Errors are Values](https://blog.golang.org/errors-are-values)


### Performance

[Profiling Go](http://www.integralist.co.uk/posts/profiling-go/)


### Concurrency

[Go memory model](https://golang.org/ref/mem)
[Additional Go tools](https://godoc.org/golang.org/x/tools)

### Advanced Object-Oriented Design

[Solid GO Design](https://dave.cheney.net/2016/08/20/solid-go-design)


### Testing

[Let's Go - Testing Golang Programs](https://code.tutsplus.com/tutorials/lets-go-testing-golang-programs--cms-26499)


### Writing GRPC-based Microservices in Go

[Generate .proto files from Go source code](https://github.com/src-d/proteus)
### Welcome to the Future

Go is opinianated. So am I. 

- Warts	
	- No file scope names (except imports)
		Python Zen - Namespaces are 
	- Package names are resolved by the last component
	- Duck Typing (implicit implementation of interfaces)
	- Visibility rules force unit tests in package -> test code mixed with production code 
	- No generics (yet)
	- [Interface embedding can lead to runtime panic if interface methods not implemented](https://code.tutsplus.com/tutorials/lets-go-object-oriented-programming-in-golang--cms-26540)
	
	- [X doesn't implement Y method with pointer rceiver](https://stackoverflow.com/questions/40823315/go-x-does-not-implement-y-method-has-a-pointer-receiver)
	
[Should Go 2.0 support Generics](https://dave.cheney.net/2017/07/22/should-go-2-0-support-generics)

[Simplicity Debt Redux](https://dave.cheney.net)
