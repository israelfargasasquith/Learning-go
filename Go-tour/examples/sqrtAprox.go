package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z:=1.0
	for i:=0;i<10;i++{
		z-=(z*z-x)/(2*z)
	}

	return z
}

func main() {
	mine := Sqrt(2)
	std :=math.Sqrt(2)
	fmt.Println("My approximation: ",mine)
	fmt.Println("STD approximation: ",std)
	fmt.Println("Diff: ",std-mine)
	
}
