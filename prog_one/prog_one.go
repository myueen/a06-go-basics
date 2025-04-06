package main

import "fmt"

func power(n int, P int) {
	res := 1

	// helper three-parameter function
	go power_h(n, P, res)
}

func power_h(n int, p int, res int) {
	if p == 1 {
		fmt.Printf("This is the result %d \n", res)
	} else {
		res = res * n
		go power_h(n, p-1, res)
	}
}

func main() {
	n := 127
	p := 12

	// start a process
	go power(n, p)
}
