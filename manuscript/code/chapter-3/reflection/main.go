package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	Number int
	Text   string
}

func main() {
	f := Foo{Number: 4, Text: "Four"}

	m := map[int]*Foo{77: &f, 88: &f}
	mv := reflect.ValueOf(m)
	fmt.Println("type of m:", mv.Type())
	fmt.Println("kind of m:", mv.Kind())
	fmt.Println("key type of m:", mv.Type().Key())

	// Get the Foo pointer value
	elem := reflect.ValueOf(m[77])
	fmt.Println("element:", elem)

	// Dereference the Foo pointer to a Foo value
	foo := elem.Elem()
	for i := 0; i < foo.NumField(); i++ {
		fmt.Println("Field", i, foo.Field(i))
	}
}
