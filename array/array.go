package main

func Sum(x []int) (sum int) {
	for _, v := range x {
		sum += v
	}
	return
}

func SumAll(x ...[]int) []int {
	var sum []int
	for _, v := range x {
		sum = append(sum, Sum(v))
	}
	return sum
}

func SumAllTails(x ...[]int) []int {
	var tail []int
	for _, v := range x {
		if len(v) == 0 {
			tail = append(tail, 0)
			continue
		}
		tail = append(tail, Sum(v[1:]))

	}
	return tail
}
