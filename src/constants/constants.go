package constants

import "fmt"

// Constants
const (
	Pi    = 3.14
	Truth = false
	Big   = 1 << 62
	Small = Big >> 61
)

// Constants Print contants
func Constants() {
	const Greeting = ""
	fmt.Println(Pi)
	fmt.Println(Truth)
	fmt.Println(Big)
	fmt.Println(Small)
}
