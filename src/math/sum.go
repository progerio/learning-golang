package math

// Sum Method Sum list of numbers
func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
