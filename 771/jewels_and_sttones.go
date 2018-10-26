package main

func numJewelsInStones(J string, S string) int {
	stones := make(map[rune]int, 0)

	for _, stone := range S {
		stones[stone] += 1
	}

	count := 0
	for _, jewel := range J {
		if num, ok := stones[jewel]; ok {
			count += num
		}
	}

	return count
}

func main() {

}
