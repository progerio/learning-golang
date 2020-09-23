package structs

// Artist Struct of Artist
type Artist struct {
	Name, Genre string
	Songs       int
}

// NewRelease Add new Songs
func NewRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}
