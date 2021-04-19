package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	myFl, _ := os.Open("/Users/AEdapa/Personal/justhackit_github/learn_go/.gitignore")
	theFile := bufio.NewScanner(myFl)

	for theFile.Scan() {
		fmt.Printf("%s\n", theFile.Text())
	}
}
