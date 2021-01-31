package main

import (
	"fmt"
	"os"
)

func main() {
	var arr = [3]int{4, 5, 6}

	for i, val := range arr {
		fmt.Fprintf(os.Stdout, "Value at %d is %d\n", i, val)
	}

}
