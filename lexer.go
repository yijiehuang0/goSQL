package main

import "fmt"

/*
	lexer.go finds every distinct group of characters in source code: tokens

 */
type keyword string
type tokenKind uint
type symbol string

/*
	Keyword operations that are supported
 */
const (
	selectKeyword keyword = "select"
	fromKeyword keyword = "from"
	asKeyword keyword = "as"
	tableKeyword keyword = "table"
	createKeyword keyword = "create"
	insertKeyword keyword = "insert"
	intoKeyword keyword = "into"
	valuesKeyword keyword = "values"
	intKeyword  keyword = "int"
	textKeyword keyword = "text"
)


/*
	Symbol operations that are supported
 */

const (
	semicolonSymbol symbol = ";"
	asteriskSymbol symbol = "*"
	commaSymbol symbol = ","
	leftParentheses symbol = "("
	rightParentheses symbol = ")"
)


/*
	Token operations that are supported
*/
const (
	keywordKind tokenKind = iota
	symbolKind
	identifierKind
	stringKind
	numericKind
)

type location struct {
	line uint
	col uint
}

type token struct {
	value string
	kind tokenKind
	loc location
}

type cursor struct {
	pointer uint
	loc location
}

func (t *token) equals(other *token) bool {
	return t.value == other.value && t.kind == other.kind
}

// defines the main function for the lexer
type lexer func(string, cursor) (*token, cursor, bool)

// we are going to return a pointer to a list of tokens in this case
// and we are going to take in a string
func lex(source string) ([]*token, error) {
	tokens := []*token{}
	cur := cursor{}

lex:
	for cur.pointer < uint(len(source)) { // while it is the case that the cur pointer is less than the len(soruce)
		lexers := []lexer{lexKeyword, lexSymbol, lexString, lexNumeric, lexIdentifier} // we are going to define a list of funcs that we are gong to range through
		for _, l := range lexers { // iterate through the functions
			if token, newCursor, ok := l(source, cur); ok {
				cur = newCursor

				// Omit nil tokens for valid, but empty syntax like newlines
				if token != nil {
					tokens = append(tokens, token)
				}

				continue lex
			}
		}

		hint := ""
		if len(tokens) > 0 {
			hint = " after " + tokens[len(tokens)-1].value
		}
		return nil, fmt.Errorf("Unable to lex token%s, at %d:%d", hint, cur.loc.line, cur.loc.col)
	}

	return tokens, nil
}
