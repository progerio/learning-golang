package functions

func Add(x , y int) int {
	return x + y
}

func Location(city string) ( string , string) {
	var region string
	var continent string 
	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown" , "Unknown" 
	}
	return region, continent
}
