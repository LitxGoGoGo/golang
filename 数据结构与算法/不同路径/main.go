package main

import (
	"fmt"
	"math/big"
)

func main() {
	m, n := 4, 2
	v1 := testUnique(m, n)
	fmt.Println(v1)
}

func uniquePaths(m, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}

func testUnique(m, n int) int {
	return int(new(big.Int).Binomial(int64(m), int64(n)).Int64())
}
