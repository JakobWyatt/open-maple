package main

import (
	"bufio"
	"fmt"
	"os"
)

//simple example showing use of tokenizer
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("There was an input error.")
		} else {
			fmt.Println(tokenizer(line))
		}
	}
}
