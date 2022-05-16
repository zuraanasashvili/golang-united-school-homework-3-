package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, len(input))
	i := 0
	for k := range input {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	var res []string
	for j := range keys {
		res = append(res, input[j])
	}
	return res
}
