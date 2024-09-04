package functions

func Calculate(aray []int, operation string) int {
	res := 0

	for _, val := range aray {
		if operation == "+" {
			res += val
		}
		if operation == "-" {
			res -= val
		}
	}

	return res
}
