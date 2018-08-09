package main

import (
	"errors"
)

//checkTokens checks an array of tokens, and verifies that:
//	-There are no undecided tokens
func checkTokens(tokens []token) error {
	for _, tokenVal := range tokens {
		if tokenVal.group == nullToken {
			errorMessage := "This token was not recognised: "
			return errors.New(errorMessage + tokenVal.value)
		}
	}
	return nil
}

func checkRoot(root tree) error {
	var err error
	err = nil
	if len(root.nodes) == 0 {
		err = errors.New("This program does nothing, " +
			"did you forget to add a semicolon?")
	}
	return err
}

func checkAssign(root tree) error {
	var err error
	err = nil
	if len(root.nodes) < 2 {
		err = errors.New("An assignment operator needs " +
			"a left and right argument.")
	} else if len(root.nodes) > 2 {
		err = errors.New("Something has gone terribly wrong, " +
			"and an assignment operator has more than 3 arguments. " +
			"Please file an issue with the code you tried to execute on github, " +
			"because it is impressive you even managed to achieve this.")
	} else if root.nodes[0].group != variable {
		err = errors.New("The LHS of an assignment operator " +
			"can only be a name.")
	} else if root.nodes[1].group == assign {
		err = errors.New("The RHS of an assignment operator " +
			"cannot be another assignment operator.")
	} else if root.nodes[1].group == rootNode {
		err = errors.New("Something has gone terribly wrong, " +
			"and an assignment operator somehow has a root node " +
			"as a subNode. Please contact the maintainers of this project.")
	}
	return err
}

func checkOperate(root tree) error {
	var err error
	err = nil
	if len(root.nodes) < 2 {
		err = errors.New("An operator needs " +
			"a left and right argument.")
	} else if len(root.nodes) > 2 {
		err = errors.New("Something has gone terribly wrong, " +
			"and an operator has more than 3 arguments. " +
			"Please file an issue with the code you tried to execute on github, " +
			"because it is impressive you even managed to achieve this.")
	} else {
		for _, val := range root.nodes {
			if val.group == assign {
				err = errors.New("The RHS and LHS of an operator " +
					"can only be a number or a variable.")
			} else if val.group == rootNode {
				err = errors.New("Something has gone terribly wrong, " +
					"and an operator somehow has a root node " +
					"as a subNode. Please contact the maintainers of this project.")
			}
		}
	}
	return err
}

func checkNumVar(root tree) error {
	var err error
	err = nil
	if len(root.nodes) != 0 {
		err = errors.New("A number or name should not have a subnode. " +
			"Please contact the maintainers of this project.")
	}
	return err
}

func checkNode(root tree) error {
	var err error
	err = nil
	switch root.group {
	case assign:
		err = checkAssign(root)
	case rootNode:
		err = checkRoot(root)
	case operate:
		err = checkOperate(root)
	case variable:
		err = checkNumVar(root)
	case constant:
		err = checkNumVar(root)
	default:
	}
	return err
}

//checkAST checks an abstract syntax tree, and verifies that:
//	-The root node has subnodes (no semicolons)
//  -Assignment nodes have correct subnodes
// 	-Operate nodes have correct subnodes
//	-Numbers and variables do not have subnodes
func checkAST(root tree) error {
	//checkAST will recursively call itself to check for errors
	err := checkNode(root)
	if err != nil {
		return err
	}
	for _, val := range root.nodes {
		//if there is an error in a subnode,
		//	we will return with this error
		err := checkAST(*val)
		if err != nil {
			return err
		}
	}

	return nil
}
