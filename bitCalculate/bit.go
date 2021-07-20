package main

import "fmt"

func main() {
	var val int32 = 32;
	fmt.Printf("%32.b\n", val)
	val1 := val << 2
	fmt.Printf("%32.b\n", val1)
	val2 := val >> 2
	fmt.Printf("%32.b\n", val2)
	val3 := val ^ val1
	fmt.Printf("%32.b\n", val3)
}


