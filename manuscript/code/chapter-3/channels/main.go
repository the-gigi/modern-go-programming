package main

import "fmt"

func send(c chan<- int, v int) {
	c <- v
}

func receive(c <-chan int) {
	v := <-c
	fmt.Println(v)
}

func main() {
	c := make(chan int)

	go send(c, 1)
	go send(c, 2)
	go send(c, 4)
	go send(c, 8)
	receive(c)
	receive(c)
	receive(c)
	go send(c, 16)
	receive(c)
	receive(c)
	go send(c, 32)
	receive(c)
}
