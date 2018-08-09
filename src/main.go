package main

import "fmt"

func main() {
	tokens, err := tokenizer("6*+/4 65;")
	printTokens(tokens)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		root, err := astBuilder(tokens)
		if err != nil {
			fmt.Println(err.Error())
		}
		treeLook(root)
	}
}
