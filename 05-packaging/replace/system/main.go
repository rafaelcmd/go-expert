package main

import (
	"fmt"
	"github.com/rafaelcmd/go-expert/05-packaging/lib"
)

func main() {
	sum := lib.Math{A: 1, B: 2}
	fmt.Println(sum.Add())
}
