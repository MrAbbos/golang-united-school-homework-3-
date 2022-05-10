package homework

import "fmt"

func sortMapValues(input map[int]string) (result []string) {

	fmt.Println(input)
	for i := 0; i <= len(input); i++ {
		fmt.Println(input[i], i)
		if _, ok := input[i]; ok {
			result = append(result, input[i])
		}
	}
	return
}
