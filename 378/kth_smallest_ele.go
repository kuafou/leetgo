package main

import "fmt"

func kthSmallest(matrix [][]int, k int) int {
	if k == 1 {
		return matrix[0][0]
	}
	l := len(matrix)
	if k == l*l {
		return matrix[l-1][l-1]
	}

	sortedMatrix := make([]int, l*l)
	for i = 0; i < l; i++ {
		for j = 0; j < l; j++ {

		}
	}

}

func main() {
	fmt.Println("vim-go")
}
