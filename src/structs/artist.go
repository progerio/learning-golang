package structs

type Artist struct {
	Name, Genre string
	Songs       int
}

func NewRelease( a Artist) int  {
	a.Songs++
	return a.Songs
}
