package main

import (
	"io/ioutil"
)

func main() {
	readFromAFile()
}

func readFromAFile() {
	data, _ := ioutil.ReadFile("/Users/aedapa/Personal/justhackit_github/learn_go/learn/hello_world.go")
	print(string(data))
}
