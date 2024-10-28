package main

import (
	"fmt"
	"github.com/rafaelcmd/go-expert/05-packaging/1/math"
)

func main() {
	sum := math.Math{A: 1, B: 2}
	fmt.Println(sum.Add())
}
