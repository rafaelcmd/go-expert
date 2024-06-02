package main

func main() {
	nums := []int{1, 2, 3, 4, 5}
	total := sum(nums...)
	println(total)
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
