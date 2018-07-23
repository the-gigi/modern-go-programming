package main

import "fmt"

func main() {
	//var ints [4]int
	//var pies = [2]float64 {3.141592, math.Pi}
	//bools := [6]bool{true, false, true}
	//fmt.Println("ints:", ints)
	//fmt.Println("pies:", pies)
	//fmt.Println("bools:", bools)
	//
	//multi := [2][4]int{}
	//for i := 0; i < 2; i++ {
	//	for j := 0; j < 4; j++ {
	//		multi[i][j] = i + j
	//}
	//}
	//
	//fmt.Println(multi)

	//x := [...]int8{1, 2, 3, 4, 5}
	//fmt.Println("x:", x, "length:", len(x), "capacity:", cap(x))

	//s := []int{1,2,3}
	//fmt.Println("s:", s, "length:", len(s), "capacity:", cap(s))
	//s = append(s, 4)
	//fmt.Println("s:", s, "length:", len(s), "capacity:", cap(s))
	//
	//fmt.Println(y[12])
	//fmt.Println(y[cap(y) + 7])

	//a2 := [2]string{"a", "b"}
	//a3 := [3]string{"c", "d", "e"}
	//
	//s2 := a2[:]
	//s3 := a3[:]
	//
	//fmt.Println(s2)
	//s2 = s3
	//fmt.Println(s2)

	//bytes := make([]byte, 100)
	//fmt.Println("length:", len(bytes), "capacity:", cap(bytes))
	//
	//
	//bytes = bytes[0:0]
	//fmt.Println("length:", len(bytes), "capacity:", cap(bytes))

	//bytes := []byte{}
	//bytes = append(bytes, 1) // append one element
	//fmt.Println("bytes", bytes, len(bytes), "capacity:", cap(bytes))
	//bytes = append(bytes, 2, 3, 4) // append multiple elements
	//fmt.Println("bytes", bytes, len(bytes), "capacity:", cap(bytes))
	//bytes = append(bytes, []byte{5,6}...) // append another slice
	//fmt.Println("bytes", bytes, len(bytes), "capacity:", cap(bytes))
	//bytes = append(bytes, bytes...) // append the slice to itself
	//fmt.Println("bytes", bytes, len(bytes), "capacity:", cap(bytes))

	//fib := []byte{0, 1, 1, 2, 3, 5, 0, 0, 0, 0, 0}
	//fmt.Println("fib", fib, len(fib), "capacity:", cap(fib))
	//fib2 := []byte{8, 13, 21, 34, 55}
	//fmt.Println("fib", fib2, len(fib2), "capacity:", cap(fib2))
	//copy(fib[6:8], fib2) // copy only two elements
	//fmt.Println("fib", fib, len(fib), "capacity:", cap(fib))
	//
	//copy(fib[6:], fib2) // copy all elements
	//fmt.Println("fib", fib, len(fib), "capacity:", cap(fib))
	//
	//fib3 := fib[6:8]
	//fib3[0] = 77 // modifies the original backing array of fib
	//fmt.Println("fib:", fib)
	//fmt.Println("fib3:", fib3)
	//
	//fib3 = append(fib3, 88) // fib3 now has a new array
	//fmt.Println("fib:", fib)
	//fmt.Println("fib3:", fib3)

	//a := [...]byte{1,2,3,4,5}
	//s := a[:]
	//fmt.Println("a", a, len(a), "capacity:", cap(a))
	//fmt.Println("s", s, len(s), "capacity:", cap(s))
	//s = append(a[:1], a[2:]...)
	//fmt.Println("a", a, len(a), "capacity:", cap(a))
	//fmt.Println("s", s, len(s), "capacity:", cap(s))

	a := [...]byte{1, 2, 3, 4, 5}
	s := a[:]
	fmt.Println("a", a, len(a), "capacity:", cap(a))
	fmt.Println("s", s, len(s), "capacity:", cap(s))
	s[1] = s[len(s)-1]
	s = s[:len(s)-1]
	fmt.Println("a", a, len(a), "capacity:", cap(a))
	fmt.Println("s", s, len(s), "capacity:", cap(s))

}
