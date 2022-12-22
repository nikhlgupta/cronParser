package main

import (
	"cronParser/expression_parser"
	"cronParser/output_formatter"
	"fmt"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 1 {
		out, err := expression_parser.ParseCronExpression(argsWithoutProg[0])
		if err != nil {
			fmt.Println("error"  + err.Error())
		}
		output_formatter.FormatOutput(out)
	} else {
		fmt.Println("invalid input to program")
	}
}