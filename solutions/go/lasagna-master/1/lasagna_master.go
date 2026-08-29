package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minsPerLayer int) int {
	if minsPerLayer == 0 {
		minsPerLayer = 2
	}

	return len(layers) * minsPerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodlesCount := 0
	sauceCount := 0

	for _, l := range layers {
		if l == "noodles" {
			noodlesCount += 1
		}

		if l == "sauce" {
			sauceCount += 1
		}
	}

	return noodlesCount * 50, float64(sauceCount) * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	result := make([]float64, 0, len(quantities))

	for _, q := range quantities {
		result = append(result, q*(float64(portions)/2))
	}

	return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
