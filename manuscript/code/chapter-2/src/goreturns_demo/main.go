package main

import (
	"fmt"
	"errors"
)

func fail() (string, int, error) {
 return errors.New("FAIL!")
}

func main() {
	fail()}
