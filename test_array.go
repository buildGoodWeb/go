package main

import "fmt"

/*
var variable_name [size] variable_type

var balance = [5] float32 {1000.0, 2.0, 3.4, 7.0, 50.0}
 */

func main() {
	var n [10]int
	var i, j int

	/* initialize elements of array n to 0 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	/* output each array element's value */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}