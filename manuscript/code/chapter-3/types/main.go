package main

import "fmt"

//type A struct {
//a int
//b string
//c bool
//}
//
//type B = A
//
//
//func main() {
//	a := A{a: 1, b: "two", c: false}
//	b := B{a: 3, b: "six", c: true}
//
//	b = a
//
//	fmt.Println(b)
//}

//func foo(v interface{}) {
//	fmt.Println(v.(int))
//}
//
//
//func main() {
//	foo("a string")
//}

func foo(v interface{}) {
	ii, ok := v.(int)
	if ok {
		fmt.Println(ii)
		return
	}

	ss, ok := v.(string)
	if ok {
		fmt.Println(ss)
		return
	}

	fmt.Println("Unknown type for value:", v)

}

func main() {
	foo(4)
	foo("a string")
	foo([]bool{true, false})
}
