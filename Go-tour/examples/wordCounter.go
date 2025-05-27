package main

import (
	"fmt"
	"strings"
	"unicode"
)

var text string = "Somebody once told me the world is gonna roll me I ain't the sharpest tool in the shed She was looking kind of dumb With her finger and her thumb"

func main() {
	f := func(c rune) bool {
		return !unicode.IsDigit(c)&&!unicode.IsLetter(c)&&unicode.IsSpace(c)
	}
	var strings[]string = strings.FieldsFunc(text , f)
	var mapWC = make(map[string]int)
	for _,v := range strings{
		mapWC[v] = mapWC[v]+1
	} 
	fmt.Println("Fields are: %q",strings )
	fmt.Println("The map lenght is:",len(mapWC))
	fmt.Println("The map values are:",mapWC)
}
