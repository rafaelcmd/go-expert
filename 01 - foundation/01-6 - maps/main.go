package main

import "fmt"

func main() {
	salaries := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	fmt.Println(salaries["steve"])
	delete(salaries, "jamie")
	salaries["mike"] = 9000

	for k, v := range salaries {
		fmt.Println(k, v)
	}

	salaries1 := make(map[string]int)
	salaries1["john"] = 10000

	for k, v := range salaries1 {
		fmt.Println(k, v)
	}

	salaries2 := map[string]int{}
	salaries2["mary"] = 15000

	for k, v := range salaries2 {
		fmt.Println(k, v)
	}
}
