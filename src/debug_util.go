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

//treeEqual tests if two trees are exactly equal,
//	including equality of all subnodes
//Does not test if the nodes have the same memory address.
func treeEqual(root1 tree, root2 tree) bool {
	//check if the root nodes are equal
	if root1.value != root2.value || root1.group != root2.group {
		return false
	}
	//check that we can iterate evenly over all subNodes
	if len(root1.nodes) != len(root2.nodes) {
		return false
	}

	for i := range root1.nodes {
		if !treeEqual(*root1.nodes[i], *root2.nodes[i]) {
			return false
		}
	}

	return true
}
