

## Why Go?

https://blog.golang.org/8years

## Reference

[The docs](https://golang.org/doc/)

[The spec](https://golang.org/ref/spec)

[A Taste of Go](https://talks.golang.org/2014/taste.slide#1)

[Language Design in the service of Software Engineering](https://talks.golang.org/2012/splash.article#TOC_17.)

[The Evolution of Go](https://talks.golang.org/2015/gophercon-goevolution.slide#1)

### Type System

[Understanding the Go Type System](https://thenewstack.io/understanding-golang-type-system/)


### Get Up and Running

[Go command](https://golang.org/cmd/go/)

[dep]()https://github.com/golang/dep)


### Error Handling

[Errors are Values](https://blog.golang.org/errors-are-values)
[Error Handling in Go](https://scene-si.org/2017/11/13/error-handling-in-go/)


### Performance

[Profiling Go](http://www.integralist.co.uk/posts/profiling-go/)
https://github.com/dgryski/go-perfbook

### Concurrency

[Go memory model](https://golang.org/ref/mem)
[Additional Go tools](https://godoc.org/golang.org/x/tools)

### Advanced Object-Oriented Design

[Solid GO Design](https://dave.cheney.net/2016/08/20/solid-go-design)


### Testing

[Let's Go - Testing Golang Programs](https://code.tutsplus.com/tutorials/lets-go-testing-golang-programs--cms-26499)

### Perofrmance

https://rakyll.org/pprof-ui/
https://github.com/dgryski/go-perfbook/blob/master/performance.md


### Writing Command-Line Tools in Go

[termui](https://github.com/gizak/termui)


## gRPC

[Generate .proto files from Go source code](https://github.com/src-d/proteus)

[GO gRPC - Beyond the Basics](https://blog.gopheracademy.com/advent-2017/go-grpc-beyond-basics/)

[Twirp - Alternative to gRPC](https://blog.twitch.tv/twirp-a-sweet-new-rpc-framework-for-go-5f2febbf35f)

### Welcome to the Future

Go is opinianated. So am I. 

# Warts	
    - [Tons of useful gothas](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)
	- No file scope names (except imports)
		Python Zen - Namespaces are 
	- Package names are resolved by the last component
	- Duck Typing (implicit implementation of interfaces)
	- Visibility rules force unit tests in package -> test code mixed with production code 
	- No generics (yet)
	- [Interface embedding can lead to runtime panic if interface methods not implemented](https://code.tutsplus.com/tutorials/lets-go-object-oriented-programming-in-golang--cms-26540)
	
	- [X doesn't implement Y method with pointer rceiver](https://stackoverflow.com/questions/40823315/go-x-does-not-implement-y-method-has-a-pointer-receiver)
	
	- time.Format is messed up: "01-02-2006 15:04:06 -07:00"
	
	- enum/constants/iota is messed up, especially json serialization
	
	- error shadowing requires var declaration inside scopes when err is declared outside.
	
## Enum json serialization

```
package main

import (
    "encoding/json"
    "fmt"
)

type EnumType int

const (
    Zero     EnumType = iota
    One
)


type EnumContainer struct {
    Name                string    `json:"name"`
    Value               EnumType `json:"value"`
}

func (e *EnumType) UnmarshalJSON(data []byte) error {
    var s string
    err := json.Unmarshal(data, &s)
    if err != nil {
        return err
    }

    value, ok := map[string]EnumType{"Zero": Zero, "One": One}[s]
    if !ok {
        return errors.New("Invalid EnumType value")
    }
    *e = value
    return nil

}

func (e *EnumType) MarshalJSON() ([]byte, error) {
    value, ok := map[EnumType]string{Zero: "Zero", One:"One"}[*e]
    if !ok {
        return errors.New("Invalid EnumType value")
    }
    return json.Marshal(value)
}


func main() {
    x := One
    ec := EnumContainer{
        "1111",
        &x,
    }
    s, err := json.Marshal(ec)
    if err != nil {
        fmt.Printf("fail!")
    }

    var ec2 EnumContainer
    err = json.Unmarshal(s, &ec2)

    fmt.Println(ec2)
    fmt.Println(*ec2.Value)
}
```
	
	
[Should Go 2.0 support Generics](https://dave.cheney.net/2017/07/22/should-go-2-0-support-generics)

[Simplicity Debt Redux](https://dave.cheney.net)


## Chapters

### Chapter 1

### Chapter 2

https://blog.golang.org/godoc-documenting-go-code

https://blog.golang.org/introducing-gofix

https://golang.org/doc/code.html
https://code.tutsplus.com/tutorials/lets-go-golang-code-organization--cms-27368
https://talks.golang.org/2014/organizeio.slide
https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2

[How I use GOPATH with multiple workspaces](https://dmitri.shuralyov.com/blog/18)

### Chapter 3


### Chapter 4


### Chapter 5


### Chapter 6


### Chapter 7


### Chapter 8


### Chapter 9


### Chapter 10

https://crawshaw.io/blog/tragedy-of-finalizers
https://crawshaw.io/blog/sharp-edged-finalizers

### Chapter 11


### Chapter 12


### Chapter 13


### Chapter 14


### Chapter 15


### Chapter 16


### Chapter 17


### Chapter 18


### Chapter 19


### Chapter 20


### Chapter 21
