package main

import "fmt"

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if s2ContainsLetterThatS1NotHas(s1, s2) {
		return 0
	}

	//

}

func s2ContainsLetterThatS1NotHas(s1, s2) bool {
	letterMap := make(map[rune]struct{}, 0)
	for _, l := range s1 {
		letterMap[l] = struct{}{}
	}
	for _, l := range s2 {
		if _, ok := letterMap[l]; !ok {
			return true
		}
	}

	return false
}

func main() {
	count := getMaxRepetitions("acb", 4, "ab", 2)
	fmt.Println(count)
}
