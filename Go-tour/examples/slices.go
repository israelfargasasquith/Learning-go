package slices

import "fmt"

func main() {
	s := []int{1,2,3}
	s[1]=4
	printSlice(s)
	var s2 []int = s[0:3]
	s2[0] = 99
	printSlice(s2)
	s2 = append(s2, s2[0],s2[1],s2[2])
	printSlice(s2)
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
