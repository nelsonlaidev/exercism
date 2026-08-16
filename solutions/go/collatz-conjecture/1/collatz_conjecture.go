package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("the number must be positive")
	}

	stepCount := 0

	for {
		if n == 1 {
			break
		}

		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}

		stepCount += 1
	}

	return stepCount, nil
}
