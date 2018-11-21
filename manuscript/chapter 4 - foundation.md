# Foundations
	
This chapter means business. I kid you not! You will get to see all the building blocks of programming, see where Go stands on the shoulders of giants and where it decided to add its own variation on a theme. The topics we will cover are:
 
 - Naming of different entities such as files, types and variables
 - How does assignment work in Go
 - The Coolness that is Type Inference
 - Controlling the flow of your program with conditions and switching
 - Looping and Iterations
	
## Naming and Visibility

Go takes the stance that identifiers starting with a Capital letter are public (exported) and identifiers starting with a lowercase letter are private. That rule applies to constants, variables, types, functions, interfaces and what have you. The visibility boundary is Go is the package. Any module in a given package can access any private and public entities defined in that package. Any module outside the package can only access public entities of the package. In addition, there is also local scope. Any identifier declared inside a function can be accessed in this function only. Identifiers declared inside inner scopes can be accessed inside that scope and its nested scopes.

Identifiers in Go can use any UTF-8 letter or digit and must start with a letter. However, public identifiers must start with an English capital letter. This breaks somewhat the internationality of Go, since English gets this privileged status regarding public identifiers.

This snippet demonstrates all these concepts:


```
package main

import "fmt"

var PublicVar = "I am public"
var privateVar = "I am private"

func PublicFunc() {
	fmt.Println(PublicVar)
}

func privateFunc() {
	fmt.Println(privateVar)
}

func main() {
	PublicFunc()
	privateFunc()

	localVar := "I am local"
	localFunc := func() {
		fmt.Println(localVar)
	}

	{
		localFunc()
		nestedScopeVar := "I am nested"
		fmt.Println(nestedScopeVar)

	}
}
```

Note that the main() function is always private. Code from another package is not supposed to call main(). It is invoked automatically when running the Go program that includes it. You can call main() directly from inside the main package if you want to. This can be useful for testing.


## Variables, Pointers and Assignment

## Type Inference

## Control Flow

### If-else

### No ternary operator

### Switch

### The Mighty For Loop

## Iterations

### Iterating Over Arrays and Slices

### Iterating Over Maps

### Iterating Over Channels
