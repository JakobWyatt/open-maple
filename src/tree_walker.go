package main

import (
	"errors"
	"strconv"
)

//binaryEvaluate takes a root node and applies the given function to each subnode
func binaryEvaluate(root tree, symbolTable map[string]interface{}, fn binaryFunc) (interface{}, error) {
	var err error
	err = nil
	var leftValue, rightValue, nodeValue interface{}

	//reuse error value by aborting processing as soon as it is no longer nil
	//the value of symbolTable is discarded, as it should not be mutated by the call to processNode
	leftValue, _, err = processNode(*root.nodes[0], symbolTable)
	if err == nil {
		rightValue, _, err = processNode(*root.nodes[1], symbolTable)
		if err == nil {
			nodeValue = fn(leftValue.(float64), rightValue.(float64))
		}
	}

	return nodeValue, err
}

//processOperator returns the value of an operator subNode,
//	after recursively evaluating its subnodes and applying the operation to it
func processOperator(root tree, symbolTable map[string]interface{}) (interface{}, error) {
	var err error
	err = nil
	var nodeValue interface{}

	switch root.value {
	case "+":
		nodeValue, err = binaryEvaluate(root, symbolTable, add)
	case "-":
		nodeValue, err = binaryEvaluate(root, symbolTable, subtract)
	case "*":
		nodeValue, err = binaryEvaluate(root, symbolTable, multiply)
	case "/":
		nodeValue, err = binaryEvaluate(root, symbolTable, divide)
	}

	return nodeValue, err
}

//processNode "evaluates" a node, by deciding the operation to be performed
//	and calling itself recursively
func processNode(root tree, symbolTable map[string]interface{}) (interface{}, map[string]interface{}, error) {
	var err error
	err = nil
	var nodeValue interface{}

	switch root.group {
	case variable:
		var ok bool
		nodeValue, ok = symbolTable[root.value]
		if ok == false {
			err = errors.New("Could not find variable " + root.value)
		}
	case constant:
		nodeValue, err = strconv.ParseFloat(root.value, 64)
	case assign:
		//symbolTable should not be mutated by the call to processNode here
		//get the name of the LH node, and assign the value of the RH node to it
		symbolTable[root.nodes[0].value], _, err = processNode(*root.nodes[1], symbolTable)
	case operate:
		nodeValue, err = processOperator(root, symbolTable)
	}

	return nodeValue, symbolTable, err
}

//treeWalker walks an AST, initializing and performing operations on variables.
//It will return a map of variable names and their values,
//	after all operations in the tree have been performed.
func treeWalker(root tree, symbolTable map[string]interface{}) (map[string]interface{}, error) {
	for _, val := range root.nodes {
		_, symbolTable, err := processNode(*val, symbolTable)
		if err != nil {
			return symbolTable, err
		}
	}
	return symbolTable, nil
}
