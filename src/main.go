package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("../samples/test_samples/tokenizer_test.txt")
	if err != nil {
		fmt.Println("There was a problem reading the file")
	} else {
		fmt.Println(tokenizer(string(file)))
	}
}
