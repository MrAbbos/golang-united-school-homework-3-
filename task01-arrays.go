package homework

func average(input [15]float32) (result float32) {
	var avg float32
	for _, v := range input {
		avg += v
	}
	return avg / float32(len(input))
}
