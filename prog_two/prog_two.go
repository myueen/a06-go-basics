package main

import (
	"fmt"
	"time"
)

var res int

func power(n int, p int) {
	res = res * n
}

func main() {
	n := 127
	p := 3
	res := 1

	// start a process
	for i := 0; i < p; i++ {
		go power(n, p)
		time.Sleep(3100)
	}

	fmt.Printf("This is the result %d \n", res)
}
