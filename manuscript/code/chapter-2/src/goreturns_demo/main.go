package main

import (
	"errors"
	"fmt"
)

func fail() (string, int, error) {
	return errors.New("FAIL!")
}

func main() {
	fail()
}
