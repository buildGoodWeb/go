package main

import "fmt"

func main() {
	/* local variable definition */
	var a int = 100
	var b int = 200
	var ret int

	/* calling returning the max between two numbers */
	ret = max(a, b)

	fmt.Printf("max value is : %d\n", ret)
}

/* function returning the max between two numbers */
func max(num1, num2 int) int {
	/* local variable declaration */
	var result int

	if(num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}