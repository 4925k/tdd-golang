package integer

// Add returns the sum of given numbers
func Add(x ...int) int {
	l := len(x)
	if l == 0 {
		return 0
	}

	if l == 1 {
		return x[0]
	}

	var sum int
	for _, v := range x {
		sum += v
	}

	return sum
}
