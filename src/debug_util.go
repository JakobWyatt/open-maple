package main

import (
	"fmt"
	"strconv"

	"github.com/golang-collections/collections/stack"
)

//printNodes helps the user debug an AST
//given a node, it will print information about that node,
//	and any subnodes
func printNodes(root tree) {
	fmt.Printf("Group: %s    Value: %s\n\n", root.group.String(), root.value)
	for _, val := range root.nodes {
		fmt.Printf("Group: %s    Value: %s\n", val.group.String(), val.value)
	}
}

func printTokens(tokens []token) {
	for _, val := range tokens {
		fmt.Printf("Group: %s    Value: %s\n", val.group.String(), val.value)
	}
}

//treeLook helps the user debug an AST,
//	and provides an interface in which to analyze an AST
func treeLook(root tree) {
	//we use a stack to store the nodes above us in the tree
	println("Enter ln to list nodes,",
		"and an integer n to move to the nth node.",
		"Move up a node by entering -1,",
		"and stop looking at the tree by entering q.")

	prevNodes := stack.New()
	keepGoing := true
	for keepGoing {
		var input string
		fmt.Scan(&input)
		num, err := strconv.Atoi(input)

		switch {
		case input == "q":
			keepGoing = false
		case input == "ln":
			printNodes(root)
		//input is an integer
		case err == nil:
			if num == -1 {
				//move up a node
				if prevNodes.Len() == 0 {
					fmt.Println("This is the base node.")
				} else {
					root = prevNodes.Pop().(tree)
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
