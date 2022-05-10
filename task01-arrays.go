package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	rs := 0
	rs = float32(rs)
	for _, v := range input {
		rs += v
	}
	return rs / float(len(input))
}
