package homework

func reverse(input []int64) (result []int64) {
	var tmp_slice []int64
	for i := len(input) - 1; i >= 0; i-- {
		tmp_slice = append(tmp_slice, input[i])
	}
	return tmp_slice
}
