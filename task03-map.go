package homework

func sortMapValues(input map[int]string) (result []string) {
	
	for _, v := range input {
		result = append(result, v)
	}
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return
}
