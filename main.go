package main

import "fmt"

func main() {
	fmt.Println(findUniquePaths(4, 4, makeBoard(5, 5), makeMem(5, 5)))
}

func findUniquePaths(row int, col int, board [][]string, mem [][]int) int {
	return 1
}

func makeBoard(rows int, cols int) [][]string {
	return 1
}

func makeMem(rows int, cols int) [][]int {
	return 1
}
