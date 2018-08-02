package main

import (
	"strings"
	"unicode"
)

type tokenType int

const (
	operator tokenType = iota
	statementDelim
	name
	number
	assignment
	nullToken
)

type token struct {
	value string
	group tokenType
}

//tokenizer iterates through a string and builds a slice of tokens.
//This slice of tokens can then be built into an AST by the parser function.
func tokenizer(input string) []token {
	var tokens []token

	var tokenBuilder strings.Builder
	tokenGroup := nullToken

	pushToken := func() {
		if tokenBuilder.Len() > 0 {
			tokens = append(tokens, token{value: tokenBuilder.String(), group: tokenGroup})
			tokenBuilder.Reset()
			tokenGroup = nullToken
		}
	}

	//We iterate over a UTF-8 encoded string and tokenize
	//Each switch outcome has the choice to either add onto the current tokenBuilder,
	//	or push the tokenBuilder onto tokens and begin a new token
	//Good example of this logic is the "=" switch outcome
	for _, char := range input {
		switch {
		//A space will always end the current token
		//We will deal with undecided tokens during the error checking part of the program
		case unicode.IsSpace(char):
			pushToken()
		//A semicolon will always end the current token
		case char == ';':
			pushToken()
			tokenBuilder.WriteByte(';')
			tokenGroup = statementDelim
		//A colon will always end the current token
		//We don't know if it forms part of an assignment op or not,
		//	therefore we make its tokenGroup unknown
		case char == ':':
			pushToken()
			tokenBuilder.WriteByte(':')
		//The only current valid use for '=' is as part of an assignment operator,
		//	so if tokenBuilder does not contain only ":"
		//	we will end the current token and push a nullToken onto tokens
		case char == '=':
			if tokenBuilder.String() == ":" {
				tokenBuilder.WriteByte('=')
				tokenGroup = assignment
			} else {
				pushToken()
				tokenBuilder.WriteByte('=')
			}
		//Any operator will always end the current token
		//[-] can be the start of a number or an operator depending on context
		case char == '-':
			pushToken()
			tokenBuilder.WriteByte('-')
			if tokens[len(tokens)-1].group == operator {
				tokenGroup = number
			} else {
				tokenGroup = operator
			}
		case char == '*' || char == '/' || char == '+':
			pushToken()
			tokenBuilder.WriteByte(byte(char))
			tokenGroup = operator
		//If the current token is a name or number, the numbers will be added onto it
		//	Otherwise, we will begin a new number token
		case unicode.IsDigit(char) || char == '.':
			if !(tokenGroup == number || tokenGroup == name) {
				pushToken()
				tokenGroup = number
			}
			tokenBuilder.WriteByte(byte(char))
		//Any other character can only be a name
		default:
			if tokenGroup != name {
				pushToken()
				tokenGroup = name
			}
			tokenBuilder.WriteRune(char)
		}
	}

	return tokens
}
