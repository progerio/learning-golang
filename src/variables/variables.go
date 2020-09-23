package variables

package variables

import "fmt"

/* Can be declared
var(
	name string
	age int
	location string
)

// Or
var (
	name, location string
	age int
)

// Or
var name string
var age int
var location string

// A var declaration can include initializers, one per variable.

var (
	name string = "Prince Oberyn"
	age int = 32
	location string = "Dorne"
)

*/

var (
	name, location, age = "Prince Oberyn", "Dorne", 32
)

func Variables() {
	fmt.Println(name, location, age)
	// A variable can contain any type, including functions
	action := func() {
		fmt.Printf("Print variables from func %s, %s, %d\n", name, location, age)
	}
	action()
}
