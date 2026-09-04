package lineup

import "fmt"

func Format(name string, number int) string {
	n := ""

	switch {
	case number%10 == 1 && number%100 != 11:
		n = "st"
	case number%10 == 2 && number%100 != 12:
		n = "nd"
	case number%10 == 3 && number%100 != 13:
		n = "rd"
	default:
		n = "th"
	}

	return fmt.Sprintf("%v, you are the %v%v customer we serve today. Thank you!", name, number, n)
}
