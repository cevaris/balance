package lb

func calcWeightPercentages(weights []int) []int {
	var sum int = 0
	for _, v := range weights {
		sum += v
	}

	percents := make([]int, len(weights))
	for i, v := range weights {
		percents[i] = int((float64(v)/float64(sum)) * 100)
	}
	return percents
}
