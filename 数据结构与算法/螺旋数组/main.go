package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrix)
	spiralOrder(matrix)
}

func spiralOrder(matrix [][]int) []int {
	//螺旋填充
	var left, top, right, bottom = 0, 0, len(matrix[0]) - 1, len(matrix) - 1
	var result []int
	target := (len(matrix[0])) * (len(matrix))

	for {
		if target < 1 {
			break
		}
		for i := left; i <= right && target >= 1; i++ {
			result = append(result, matrix[top][i])
			target--
		}
		top++
		for i := top; i <= bottom && target >= 1; i++ {
			result = append(result, matrix[i][right])
			target--
		}
		right--
		for i := right; i >= left && target >= 1; i-- {
			result = append(result, matrix[bottom][i])
			target--
		}
		bottom--
		for i := bottom; i >= top && target >= 1; i-- {
			result = append(result, matrix[i][left])
			target--
		}
		left++

	}
	return result
}
