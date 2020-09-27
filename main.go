package main

import (
	"fmt"
	"time"

	"github.com/progerio/learning-golang/src/types"

	"github.com/progerio/learning-golang/src/structs"

	"github.com/progerio/learning-golang/src/functions"

	"github.com/progerio/learning-golang/src/constants"

	"github.com/progerio/learning-golang/src/variables"
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
	fmt.Printf("Location : region = %s and continent = %s\n", region, continent)
	fmt.Println("")

	fmt.Println("..........Structs..............")
	me := &structs.Artist{Name: "Paulo", Genre: "Software Enginer", Songs: 42}
	fmt.Printf("%s realases their %dth song\n", me.Name, structs.NewRelease(me))
	fmt.Printf("%s has a total of %d songs\n", me.Name, me.Songs)

	bootcamp := structs.Bootcamp{Lat: 34.012836, Lon: -118.495338, Date: time.Now()}
	fmt.Println(bootcamp)

	x := new(structs.Bootcamp)
	y := &structs.Bootcamp{}
	fmt.Println(*x == *y)
	user :=structs.User{ID: 42, Name: "Matt", Location: "LA"}
	p := structs.Player{user, 90404}
	fmt.Printf("Id: %d, Name: %s, Location: %s, Game id: %d\n", p.ID, p.Name, p.Location, p.Location)
	p.ID = 11
	fmt.Printf("%+v", p)
	fmt.Println("")

	fmt.Println("..........Types.................")
	types.PrintTypes()
	fmt.Println(".... Print types conversions....")
	types.PrintTypesConversion(42)
	fmt.Println(".... Print type assertion.......")
	foo := map[string]interface{}{
		"Matt": 42,
	}
	types.TimeMap(foo)
	fmt.Println(foo)
	fmt.Println("")

}
