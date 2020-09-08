package main

import (
	"fmt"
	"io/ioutil"

	"ingoly/utils/parser"
	"ingoly/utils/tokenizer"
)

func main() {
	data, err := ioutil.ReadFile("example.ig")

	if err != nil {
		fmt.Println(err)
	}

	lx := tokenizer.New(string(data))
	result := lx.Tokenize()
	var jp parser.Parser
	ps := jp.New(result)
	ast := ps.Parse()
	// ast.PrintRecursive()

	w := parser.NewBlockWalker()

	for _, stmt := range ast.Tree {
		stmt.Walk(w)
	}

	p := parser.NewPrinter()

	for _, stmt := range ast.Tree {
		stmt.Walk(p)
	}

	fmt.Println(w.Vars)

	// for line, instruction := range ast.Tree {
	// 	// ok := instruction.Execute()
	// 	// if ok == nil {
	// 	// 	fmt.Print("Line Num: ")
	// 	// 	fmt.Print(line)
	// 	// 	fmt.Print(" successful executed")
	// 	// 	fmt.Println()
	// 	// }
	// }

}
