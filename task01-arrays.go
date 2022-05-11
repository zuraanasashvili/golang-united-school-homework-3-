package homework

func average(input [15]float32) (result float32) {
	var rs float32 = 0
	for _, v := range input {
		rs += v
	}
	return rs / float32(len(input))
}
