package utils

func ReverseStringListMap(input map[string][]string) (output map[string][]string) {
	output = map[string][]string{}
	for key, set := range input {
		for _, value := range set {
			output[value] = append(output[value], key)
		}
	}
	return
}
