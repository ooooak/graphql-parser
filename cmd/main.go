package main

import (
	"fmt"
	"log"

	"github.com/ooooak/apj/pkg/graphqlparser"
)

func main() {
	data := graphqlparser.InitLexer("")
	tk, err := data.ReadToken()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(tk)
}
