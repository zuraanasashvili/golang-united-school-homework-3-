package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	rs := 0
	for _, v = range input {
		rs += v
	}
	return rs / 15
}
