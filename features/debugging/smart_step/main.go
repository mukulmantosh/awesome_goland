package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Concat(ToUpper("hello"), ToUpper("world")))
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func Concat(s1, s2 string) string {
	return s1 + " " + s2
}
