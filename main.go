package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	listOfTokens, error := lex(text)
	if error == nil {
		for index, token := range(listOfTokens) {
			fmt.Println(index, token.value, token.kind, token.loc.col, token.loc.line)
		}
	}


}
