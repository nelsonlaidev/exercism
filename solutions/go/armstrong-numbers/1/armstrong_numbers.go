package armstrongnumbers

import (
	"math"
	"strconv"
	"strings"
)

func IsNumber(n int) bool {
	str := strconv.Itoa(n)
	slice := strings.Split(str, "")
	sum := 0

	for _, d := range slice {
		digit, _ := strconv.Atoi(d)
		sum += int(math.Pow(float64(digit), float64(len(slice))))
	}

	return sum == n
}
