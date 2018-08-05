package main

import "fmt"

//func changeMe(p *int) {
//	*p = *p + 77
//	fmt.Println("changeMe() - &p:", &p, "p:", p, "*p:", *p)
//}
//
//
//func main() {
//	x := 4
//	p := &x
//	fmt.Println("main()     - &p:", &p, "p:", p, " x:", x)
//	changeMe(p)
//	fmt.Println("main()     - &p:", &p, "p:", p, " x:", x)
//}

func changeMe(m map[string]int) {
	m["x"] = 7
}

func main() {
	m := map[string]int{"x": 5}
	fmt.Println("main()", m)
	changeMe(m)
	fmt.Println("main()", m)
}
