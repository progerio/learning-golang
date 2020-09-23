package types

import (
	"fmt"
	"math/cmplx"
)

var (
	goIsFun bool       = true
	maxInt  uint64     = 1<<64 - 1
	complex complex128 = cmplx.Sqrt(-5 + 12i)
)

// PrintTypes Print types basics
func PrintTypes() {
	const f = "%T(%v)\n"
	fmt.Printf(f, goIsFun, goIsFun)
	fmt.Printf(f, maxInt, maxInt)
	fmt.Printf(f, complex, complex)
}
// PrintTypesConversion Print types conversions
func PrintTypesConversion(i int) {

	f := float64(i)
	u := uint(f)
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(u)
}

