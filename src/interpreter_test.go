package main

import (
	"io/ioutil"
	"strconv"
	"testing"
)

func expectedOutputInterpretTestv1() map[string]interface{} {
	expectedOutput := make(map[string]interface{})
	expectedOutput["arg01"] = 2.5
	expectedOutput["arg02"] = -32.5
	expectedOutput["arg03"] = 2.5/13 - 32.5
	expectedOutput["arg04"] = 2.5/13 - 20.5

	return expectedOutput
}

func TestInterpreter(test *testing.T) {
	testFilePath := "../samples/test_samples/interpret_test_v1/"
	maple := interpreter()

	var symbolTable map[string]interface{}
	var err error
	var code []byte

	for i := 0; i != 3; i++ {
		code, err = ioutil.ReadFile(testFilePath + strconv.Itoa(i) + ".txt")
		if err != nil {
			test.Fatal("Could not read test file.")
		}
		symbolTable, err = maple(string(code))
		if err != nil {
			test.Fatal("There was an error interpreting the program, " +
				"where there should have been none. Error: " + err.Error())
		}
	}

	if !symbolTableEqual(symbolTable, expectedOutputInterpretTestv1()) {
		test.Error("Interpreter did not return the correct symbolTable.")
	}
}
