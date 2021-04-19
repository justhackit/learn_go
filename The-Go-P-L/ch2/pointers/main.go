package main

import "fmt"

func incr(i *int) int {
	*i++
	return *i
}

func incrV1(i int) int {
	return i + 1
}

func main() {
	t := 1
	fmt.Printf("incrV1 : %d and t is now %d\n", incrV1(t), t)
	fmt.Printf("incr : %d and t is now %d\n", incr(&t), t)
}
