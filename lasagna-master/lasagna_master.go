package lasagna_master

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return timePerLayer * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {

	var noodles int = 0
	var sauce float64 = 0.0

	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		} else if v == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaledRecipe []float64
	for _, v := range quantities {
		scaledRecipe = append(scaledRecipe, v*float64(portions)/2)
	}
	return scaledRecipe
}
