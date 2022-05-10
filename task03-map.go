package homework

func sortMapValues(input map[int]string) (result []string) {
	for i := 0; i <= len(input); i++ {
		if _, ok := input[i]; ok {
			result = append(result, input[i])
		}
	}
	return
}
