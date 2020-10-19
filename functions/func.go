package main

import "fmt"

func swap(x, y, z string) (string, string, string) {
	return z, y, x
}
// Print “Hello, World!” to console
func main() {
	a, b, c := swap("hello", "world", "wonderful" )
	fmt.Println(a, b, c)
}
