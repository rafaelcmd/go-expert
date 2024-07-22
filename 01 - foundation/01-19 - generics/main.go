package main

type MyNumber int

type Number interface {
	~int | float64
}

func Sum[T Number](m map[string]T) T {
	var sum T

	for _, v := range m {
		sum += v
	}

	return sum
}

func Compare[T comparable](a T, b T) bool {
	return a == b
}

func main() {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
	m3 := map[string]MyNumber{"a": 1, "b": 2, "c": 3}

	println(Sum(m1))
	println(Sum(m2))
	println(Sum(m3))
	println(Compare(1, 1))
}
