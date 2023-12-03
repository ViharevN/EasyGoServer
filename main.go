package main

import (
	"fmt"
	"math"
)

func main() {
	v := Abs(3)
	fmt.Println(v)
}

func Abs(value float64) float64 {
	return math.Abs(value)
}
