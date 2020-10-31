package space

import "math"

// Planet is a string type
type Planet string

// Age returns an age in a given planet
func Age(sec float64, planet Planet) float64 {
	oneYearInSeconds := 1.0 / 31557600.00
	timeOnEarth := sec * oneYearInSeconds
	time := timeOnEarth

	switch planet {
	case "Mercury":
		time = timeOnEarth / 0.2408467

	case "Venus":
		time = timeOnEarth / 0.61519726

	case "Earth":
		time = timeOnEarth

	case "Mars":
		time = timeOnEarth / 1.8808158

	case "Jupiter":
		time = timeOnEarth / 11.862615

	case "Saturn":
		time = timeOnEarth / 29.447498

	case "Uranus":
		time = timeOnEarth / 84.016846

	case "Neptune":
		time = timeOnEarth / 164.79132
	}

	return toFixed(time, 2)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
