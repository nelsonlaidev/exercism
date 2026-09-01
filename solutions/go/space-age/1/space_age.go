package spaceage

type Planet string

const secondsInYear = 31557600

func Age(seconds float64, planet Planet) float64 {

	switch planet {
	case "Mercury":
		return seconds / (secondsInYear * 0.2408467)
	case "Venus":
		return seconds / (secondsInYear * 0.61519726)
	case "Earth":
		return seconds / secondsInYear
	case "Mars":
		return seconds / (secondsInYear * 1.8808158)
	case "Jupiter":
		return seconds / (secondsInYear * 11.862615)
	case "Saturn":
		return seconds / (secondsInYear * 29.447498)
	case "Uranus":
		return seconds / (secondsInYear * 84.016846)
	case "Neptune":
		return seconds / (secondsInYear * 164.79132)
	default:
		return -1
	}
}
