package raindrops

import "fmt"

// Convert is a function that checks if a number has 3, 5 and 7 as a factor
func Convert(input int) string {
	var result string = ""
	if input%3 == 0 {
		result = result + "Pling"
	}
	if input%5 == 0 {
		result = result + "Plang"
	}
	if input%7 == 0 {
		result = result + "Plong"
	}

	if input%3 != 0 && input%5 != 0 && input%7 != 0 {
		result = fmt.Sprint(input)
	}

	return result
}
