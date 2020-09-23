package main

import (
	"fmt"

	"org.progerio.com/learning/src/structs"

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
	fmt.Println("The sum from 10 + 20 = ", functions.Add(10, 20))
	region, continent := functions.Location("LA")
	fmt.Printf("Location : region = %s and continent = %s", region, continent)
	fmt.Println("")

	fmt.Println("..........Structs..............")
	me := &structs.Artist{Name: "Paulo", Genre: "Software Enginer", Songs: 42}
	fmt.Printf("%s realases their %dth song\n", me.Name, structs.NewRelease(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
	fmt.Println("")
}
