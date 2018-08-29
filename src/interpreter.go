package main

func frontEnd(code string) (tree, error) {
	var err error
	err = nil
	var root tree

	tokens, err := tokenizer(code);
	if err == nil {
		root, err = astBuilder(tokens)
	}

	return root, err
}

func run(code string, symbolTable map[string]interface{}) (map[string]interface{}, error) {
	var err error
	err = nil

	root, err := frontEnd(code)
	if err == nil {
		symbolTable, err = treeWalker(root, symbolTable)
	}
	
	return symbolTable, err
}

//interpreter is a closure which stores the state of the program
//it returns a function which, when called,
//	will execute any code passed to it
func interpreter() func(string) (map[string]interface{}, error) {
	var err error
	err = nil
	var symbolTable map[string]interface{}

	//TODO: func should not return symbolTable,
	//	however until functions are properly implemented,
	//	including the print function,
	//	it has to do this to let the user know program state.
	return func(code string) (map[string]interface{}, error) {
		symbolTable, err = run(code, symbolTable)
		return symbolTable, err
	}
}
