package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"

	"github.com/golang-collections/collections/stack"
)

//execPrint executes statements until the user enters exit,
//	where the program prints the value of all current variables
func execPrint() {
	maple := interpreter()
	var symbolTable map[string]interface{}
	var codeErr error

	reader := bufio.NewReader(os.Stdin)
	var input string
	var inputErr error

	fmt.Println("Enter expressions to evaluate them. " +
		"Enter 'exit' to print a list of all current variables, " +
		"and exit the program.")

	keepGoing := true
	for keepGoing {
		input, inputErr = reader.ReadString('\n')
		if inputErr != nil {
			keepGoing = false
		} else if input == "exit\n" {
			keepGoing = false
		} else {
			symbolTable, codeErr = maple(input)
			if codeErr != nil {
				keepGoing = false
			}
		}
	}

	if codeErr != nil {
		fmt.Println(codeErr.Error())
	}
	if inputErr != nil {
		fmt.Println(inputErr.Error())
	}

	printSymbols(symbolTable)
}

//printSymbols prints all values in a symbolTable
//	This function assumes that all values in symbolTable are float64
func printSymbols(symbolTable map[string]interface{}) {
	for key, val := range symbolTable {
		fmt.Printf("%s: %f\n", key, val.(float64))
	}
}

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

//checks if the two symbolTables are equal to eachother
//it does this by checking that each entry in the tables are equal
//	this is done via the comparison operator, however if the value is a float,
//	the function uses floatEqual to determine if the two tables are the same
func symbolTableEqual(a, b map[string]interface{}) bool {
	//if the two tables have equal size, then we can iterate over a,
	// and know that there are no 'hidden' elements in b
	if len(a) != len(b) {
		return false
	}

	//a is iterated over.
	//we then make 2 checks:
	//	if the key is in b,
	//	and if both values are equal
	for key, valA := range a {
		valB, ok := b[key]
		//is key in b?
		if ok == false {
			return false
		}
		//are the types the same?
		if reflect.TypeOf(valA) != reflect.TypeOf(valB) {
			return false
		}
		//if the type is a float, then use floatEqual
		//else, use the comparision operator
		if _, ok := valA.(float64); ok {
			if !floatEqual(valA.(float64), valB.(float64)) {
				return false
			}
		} else {
			if valA != valB {
				return false
			}
		}
	}

	return true
}

func floatEqual(a, b float64) bool {
	tol := 0.0001
	return math.Abs(a-b) < tol
}
