package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "hello world jasdjfo aosdofnaosn aonsdon"
	slice := cleanInput(input)
	fmt.Println(slice)
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
