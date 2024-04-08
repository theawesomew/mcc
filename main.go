package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type TokenType int

const (
	NUMBER TokenType = iota + 1
	ADD
	SUBTRACT
	MULTIPLY
	DIVIDE
	OPENPAREN
	CLOSEPAREN
	WHITESPACE
)

type Token struct {
	value string
	ofType TokenType
}

func tokenize (document string) []Token {
	var output []Token;

	regular_expressions := map[TokenType]*(regexp.Regexp) {
		WHITESPACE: regexp.MustCompile(`^\s+`),
		NUMBER: regexp.MustCompile(`^\d+`),
		ADD: regexp.MustCompile(`^\+`),
		SUBTRACT: regexp.MustCompile(`^\-`),
		MULTIPLY: regexp.MustCompile(`^\*`),
		DIVIDE: regexp.MustCompile(`^\/`),
		OPENPAREN: regexp.MustCompile(`^\(`),
		CLOSEPAREN: regexp.MustCompile(`^\)`),
	}

	for len(document) > 0 {
		for typ, expr := range regular_expressions {
			location := expr.FindStringIndex(document)

			if location != nil {
				if typ != WHITESPACE {
					output = append(output, Token {
						value: document[location[0]:location[1]],
						ofType: typ,
					})
				}

				document = document[location[1]:]
			}
		}	
	}

	return output
}

type Operation int

const (
	NONE Operation = iota + 1
	ADDITION
	SUBTRACTION
	MULTIPLICATION
	DIVISION
)

type ExpressionTree struct {
	operator Operation
	value string
	left *ExpressionTree
	right *ExpressionTree
}

func isOperatorToken (token Token) bool {
	switch (token.ofType) {
		case ADD:
			return true;
		case SUBTRACT:
			return true;
		case MULTIPLY:
			return true;
		case DIVIDE:
			return true;	
	}

	return false;
}

func isValid (tokens []Token) bool {
	count := 0

	for i := 0; i < len(tokens); i += 1 {
		if tokens[i].ofType == OPENPAREN {
			count += 1;
		}

		if tokens[i].ofType == CLOSEPAREN {
			if count > 0 {
				count -= 1;
			} else {
				return false;
			}
		}

		if i == len(tokens) - 1 {
			break;
		}

		if tokens[i].ofType == NUMBER && tokens[i+1].ofType == NUMBER {
			return false;
		}

		if isOperatorToken(tokens[i]) && isOperatorToken(tokens[i+1]) {
			return false;
		}

		if tokens[i].ofType == CLOSEPAREN && tokens[i].ofType == OPENPAREN {
			return false;
		}
	}

	return count == 0
}

func operatorTokenToOperation (token Token) Operation {
	switch (token.ofType) {
		case ADD:
			return ADDITION
		case SUBTRACT:
			return SUBTRACTION
		case MULTIPLY:
			return MULTIPLICATION
		case DIVIDE:
			return DIVISION
	}

	return NONE
}

func unwrapBrackets (tokens []Token) []Token {
	if len(tokens) < 3 {
		return tokens
	}

	if tokens[0].ofType != OPENPAREN || tokens[len(tokens)-1].ofType != CLOSEPAREN {
		return tokens
	}

	return tokens[1:len(tokens)-1]
}

func partitionByLeastParentheticallyEnclosedOperator (tokens []Token) (Operation, string, []Token, []Token) {
	minEnclosure := int(1e6)
	indexOfMinEnclosure := -1

	tokens = unwrapBrackets(tokens)	

	for i, val := range tokens {
		if !isOperatorToken(val) {
			continue;
		}

		enclosureCount := 0

		for j := i; j > 0; j -= 1 {
			if tokens[j].ofType == OPENPAREN {
				enclosureCount += 1
			}

			if tokens[j].ofType == CLOSEPAREN {
				enclosureCount -= 1
			}
		}

		if enclosureCount < minEnclosure {
			minEnclosure = enclosureCount
			indexOfMinEnclosure = i
		}
	}

	if indexOfMinEnclosure == -1 {
		return NONE, tokens[0].value, make([]Token, 0), make([]Token, 0)
	}

	leftTokens := tokens[0:indexOfMinEnclosure]
	rightTokens := tokens[indexOfMinEnclosure+1:]	

	return operatorTokenToOperation(tokens[indexOfMinEnclosure]), "", leftTokens, rightTokens
}

func parse (tokens []Token) *ExpressionTree {
	// two NUMBER tokens next to each other is not allowed
	// two OPERATOR tokens next to each other is not allowed	
	ast := new(ExpressionTree)

	operation, value, leftTokens, rightTokens := partitionByLeastParentheticallyEnclosedOperator(tokens)

	fmt.Printf("%v %v\n", leftTokens, rightTokens)

	ast.operator = operation
	ast.value = value

	if ast.operator != NONE {
		ast.left = parse(leftTokens)
		ast.right = parse(rightTokens)
	}

	return ast;
}

func evaluate (ast *ExpressionTree) float64 {
	fmt.Printf("%v\n", ast);

	switch (ast.operator) {
		case ADDITION:
			return evaluate(ast.left) + evaluate(ast.right)
		case SUBTRACTION:
			return evaluate(ast.left) - evaluate(ast.right)
		case MULTIPLICATION:
			return evaluate(ast.left) * evaluate(ast.right)
		case DIVISION:
			return evaluate(ast.left) / evaluate(ast.right)
		case NONE:
			result, _ := strconv.ParseFloat(ast.value, 64)
			return result
	}

	return 0.0;
}

func main () {
	document := "1 + 2"
	output := tokenize(document)

	if !isValid(output) {
		log.Fatal("Invalid input")
	}

	root := parse(output)

	fmt.Printf("%v\n", output)
	fmt.Printf("%v\n", evaluate(root))
}
