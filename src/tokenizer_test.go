package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestTokenizer(test *testing.T) {
	file, err := ioutil.ReadFile("../samples/test_samples/tokenizer_test.txt")
	if err != nil {
		test.Fatal("Could not read test file.")
	}
	tokenizerOutput, err := tokenizer(string(file))

	if err != nil {
		test.Fatal("Tokenizer detected errors where there were none.")
	}

	expectedOutput := []token{
		token{value: "quadratic_user", group: name},
		token{value: ":=", group: assignment},
		token{value: "proc", group: name},
		token{value: "(a,", group: name},
		token{value: "b,", group: name},
		token{value: "c)", group: name},
		token{value: "return", group: name},
		token{value: "[(", group: name},
		token{value: "-", group: operator},
		token{value: "b", group: name},
		token{value: "-", group: operator},
		token{value: "sqrt(b^2", group: name},
		token{value: "-", group: operator},
		token{value: "4", group: number},
		token{value: "*", group: operator},
		token{value: "a", group: name},
		token{value: "*", group: operator},
		token{value: "c))", group: name},
		token{value: "/", group: operator},
		token{value: "(2", group: name},
		token{value: "*", group: operator},
		token{value: "a)", group: name},
		token{value: ";", group: statementDelim},
		token{value: "end", group: name},
		token{value: ";", group: statementDelim}}

	if !reflect.DeepEqual(tokenizerOutput, expectedOutput) {
		test.Error("Tokenizer returned incorrect token array.")
	}
}
