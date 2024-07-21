package main

func main() {
	var1 := 10
	var2 := 20
	sum(&var1, &var2)

	println(var1)
	println(var2)
}

func sum(a, b *int) int {
	*a = 50
	*b = 50

	return *a + *b
}
