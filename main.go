package main

import (
	"fmt"

	"org.progerio.com/learning/src/functions"

	"org.progerio.com/learning/src/constants"
	"org.progerio.com/learning/src/variables"
)

func main() {
	fmt.Println("..........Variables............")
	variables.Variables()
	fmt.Println("")

	fmt.Println("..........Constants............")
	constants.Constants()
	fmt.Println("")

	fmt.Println("..........Functions............")
	fmt.Println("Soma de 10 + 20 = ", functions.Add(10, 20))
	fmt.Println("")

}
