package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for _, layer := range layers {
		if layer == "sauce" {
			sauce += 0.2
		} else if layer == "noodles" {
			noodles += 50
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string){
    secretIngredient := friendsList[len(friendsList) - 1]
    myList[len(myList) - 1] = secretIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		scaledQuantities[i] = quantities[i] * (float64(portions) / 2.0)
	}
	return scaledQuantities
}
