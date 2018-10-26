package main

import "fmt"

type PairNode struct {
	Start int
	End   int
	P     *PairNode
}

func longestValidParentheses(s string) int {
	pairs := make([]byte, len(s))

	left := rune('(')
	right := rune(')')
	var wp *PairNode
	for i, p := range s {
		if i == 0 && p == right {
			continue
		}
		if p == left {
			wp = &PairNode{Start: i, P: wp}
			continue
		}
		if p == right && wp != nil {
			pairs[wp.Start] = 1
			pairs[i] = 1
			wp = wp.P
		}
	}

	longest := 0
	started := false
	round := 0
	for _, p := range pairs {
		if p == 1 {
			started = true
			round += 1
		}
		if p == 0 && started {
			if longest < round {
				longest = round
			}
			round = 0
		}
	}
	if round > longest {
		longest = round
	}

	return longest
}

func main() {
	// p := longestValidParentheses("()(()")
	// fmt.Println(p)
	// p = longestValidParentheses("()()(()")
	// fmt.Println(p)
	p := longestValidParentheses("()")
	fmt.Println(p)
}
