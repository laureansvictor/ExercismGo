package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime <= 0 {
		avgPrepTime = 2
	}
	return len(layers) * avgPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	sauce := 0.0
	noodles := 0
	for _, item := range layers {
		if item == "sauce" {
			sauce += 0.2
		}
		if item == "noodles" {
			noodles += 50
		}
	}

	return noodles, float64(sauce)
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(fList []string, mList []string) {
	mList[len(mList)-1] = fList[len(fList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, idealPortions int) []float64 {
	scaled := make([]float64, len(quantities))
	factor := float64(idealPortions) / 2.0
	for i, v := range quantities {
		scaled[i] = v * factor
	}
	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
