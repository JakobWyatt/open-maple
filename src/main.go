package main

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/golang-collections/collections/stack"
)

//printNodes helps the user debug an AST
//given a node, it will print information about that node,
//	and any subnodes
func printNodes(root tree) {
	fmt.Printf("Group: %sValue: %s", String(root.group), root.value)
	for _, val := range root.nodes {
		fmt.Printf("Group: %sValue: %s", String(val.group), val.value)
	}
}

func printTokens(tokens []token) {
	for _, val := range tokens {
		fmt.Printf("Group: %sValue: %s\n", String(val.group), val.value)
	}
}

//treeLook helps the user debug an AST,
//	and provides an interface in which to analyze an AST
func treeLook(root tree) {
	//we use a stack to store the nodes above us in the tree
	println(`Enter ln to list nodes,
		and an integer n to move to the nth node.
		Move up a node by entering -1,
		and stop looking at the tree by entering q.`)

	prevNodes := stack.New()
	for {
		var input string
		fmt.Scan(&input)
		num, err := strconv.Atoi(input)

		switch {
		case input == "q":
			break
		case input == "ln":
			printNodes(root)
		//input is an integer
		case err == nil:
			if num == -1 {
				//move up a node
				if prevNodes.Len() == 0 {
					fmt.Println("This is the base node.")
				} else {
					root = prevNodes.Peek().(tree)
					*prevNodes = prevNodes.Pop().(stack.Stack)
				}
			} else if num >= 0 && num < len(root.nodes) {
				prevNodes.Push(root)
				root = *root.nodes[num]
			} else {
				fmt.Println("That is not a valid numerical input.")
			}
		default:
			fmt.Println("That input was not recognised.")
		}

	}
}

func main() {
	file, err := ioutil.ReadFile("../samples/test_samples/tokenizer_test.txt")
	if err != nil {
		fmt.Println("There was a problem reading the file")
	} else {
		tokens := tokenizer(string(file))
		astRoot := parser(tokens, tree{group: root})
		printTokens(tokens)
		treeLook(astRoot)
	}
}
