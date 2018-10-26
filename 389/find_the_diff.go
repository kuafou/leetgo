package main

import "fmt"

func findTheDifference(s string, t string) byte {
	ss := 0
	for _, b := range s {
		ss += int(b)
	}
	st := 0
	for _, b := range t {
		st += int(b)
	}

	return byte(st - ss)
}

func main() {
	b := findTheDifference("abcd", "abcde")
	fmt.Println(b)
}
