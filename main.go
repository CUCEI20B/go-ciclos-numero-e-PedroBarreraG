package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)
	var e float64
	e = 0
	var op float64
	var n2 float64
	var n float64
	n = 1
	for n <= x {
		op = 1
		n2 = 1
		for n2 <= n {
			op = op * n2
			n2 = n2 + 1
		}
		e = e + 1/op
		n = n + 1
	}
	fmt.Println(e)
}
