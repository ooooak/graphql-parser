package main

import (
	"fmt"
	"log"

	"github.com/ooooak/apj/pkg/graphqlparser"
)

func main() {
	data := graphqlparser.InitLexer("-1.12")
	for {
		tk, err := data.ReadToken()
		if err != nil {
			log.Fatalln(err)
		}

		if tk.Kind == graphqlparser.TK_EOF {
			fmt.Println(tk)
			break
		}

		fmt.Println(tk)
	}
}
