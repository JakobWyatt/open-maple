package main

func main() { /*
		file, err := ioutil.ReadFile("../samples/test_samples/multi-statement-assign.txt")
		if err != nil {
			fmt.Println("There was a problem reading the file")
		} else {
			tokens := tokenizer(string(file))
			astRoot := astBuilder(tokens)
			printTokens(tokens)
			treeLook(astRoot)
		}*/

	tokens := tokenizer("6**9; ")
	astRoot := astBuilder(tokens)
	printTokens(tokens)
	treeLook(astRoot)
}
