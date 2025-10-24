package lasagna

// PreparationTime calculates the total preparation time based on layers and average time per layer
func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		avgTime = 2 // default value
	}
	return len(layers) * avgTime
}

// Quantities calculates the amount of noodles and sauce needed
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// AddSecretIngredient replaces the last element in myList with the last element from friendsList
func AddSecretIngredient(friendsList, myList []string) {
	if len(friendsList) > 0 && len(myList) > 0 {
		myList[len(myList)-1] = friendsList[len(friendsList)-1]
	}
}

// ScaleRecipe scales the recipe quantities for the desired number of portions
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := make([]float64, len(quantities))
	scaleFactor := float64(portions) / 2.0 // recipe is for 2 portions

	for i, qty := range quantities {
		scaled[i] = qty * scaleFactor
	}

	return scaled
}
