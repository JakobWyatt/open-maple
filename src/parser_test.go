package main

import (
	"io/ioutil"
	"testing"
)

func expectedOutputMultiStatementAssign() tree {
	//create nodes, then build tree
	root := tree{group: rootNode}
	assign0 := tree{group: assign, value: ":="}
	assign1 := tree{group: assign, value: ":="}
	multiply0 := tree{group: operate, value: "*"}
	multiply1 := tree{group: operate, value: "*"}
	multiply2 := tree{group: operate, value: "*"}
	multiply3 := tree{group: operate, value: "*"}
	divide := tree{group: operate, value: "/"}
	add := tree{group: operate, value: "+"}
	awef0 := tree{group: variable, value: "awef"}
	awef1 := tree{group: variable, value: "awef"}
	awear := tree{group: variable, value: "awear"}
	e := tree{group: variable, value: "e"}
	num345 := tree{group: constant, value: "345"}
	num12 := tree{group: constant, value: "12"}
	num34 := tree{group: constant, value: "34"}
	numneg12 := tree{group: constant, value: "-12"}
	num5 := tree{group: constant, value: "5"}
	num6 := tree{group: constant, value: "6"}

	multiply0.nodes = append(multiply0.nodes, &num345, &num12)
	multiply1.nodes = append(multiply1.nodes, &multiply0, &e)
	assign0.nodes = append(assign0.nodes, &awef0, &multiply1)
	divide.nodes = append(divide.nodes, &num34, &numneg12)
	multiply2.nodes = append(multiply2.nodes, &divide, &num5)
	multiply3.nodes = append(multiply3.nodes, &num6, &awef1)
	add.nodes = append(add.nodes, &multiply2, &multiply3)
	assign1.nodes = append(assign1.nodes, &awear, &add)
	root.nodes = append(root.nodes, &assign0, &assign1)

	return root
}

func TestParser(test *testing.T) {
	file, err := ioutil.ReadFile("../samples/test_samples/multi-statement-assign.txt")
	if err != nil {
		test.Fatal("Could not read test file.")
	}
	tokens, err := tokenizer(string(file))
	if err != nil {
		test.Fatal("Error_check detected token errors where there were none")
	}

	parserOutput, err := astBuilder(tokens)
	if err != nil {
		test.Fatal("Error_check detected errors where there were none.")
	}

	expectedOutput := expectedOutputMultiStatementAssign()

	if !treeEqual(parserOutput, expectedOutput) {
		test.Error("Parser did not return the correct tree.")
	}
}
