package lasagna

func PreparationTime(layers []string, avg_time int) int {
	if avg_time == 0 {
		avg_time = 2
	}
	return len(layers) * avg_time
}

func Quantities(layers []string) (int, float64) {
	noodles, sauce := 0, 0.0
	for _, k := range layers {
		switch k {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
	return
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	scaled := []float64{}
	scaled = append([]float64{}, amounts...)
	for i, _ := range scaled {
		scaled[i] = (scaled[i] / 2.0) * float64(portions)
	}
	return scaled
}
