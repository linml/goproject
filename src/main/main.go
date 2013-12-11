package main

import (
	"../emryshttplib"
	"fmt"
)

func hello() {
	fmt.Print("hello\n")
}

func main() {
	emryshttplib.T()
	hello()
	fmt.Print("aaaa")
}
