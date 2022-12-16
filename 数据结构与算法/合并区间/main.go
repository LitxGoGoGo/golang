package main

import "sort"

func main() {

}

func merge(intervals [][]int) [][]int {
	//先对切片进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	pre := intervals[0]

	//遍历全部的切片
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if pre[1] < cur[0] { //意味着没有重合,这个时候就可以往结果res中推入pre了，并且这个时候的cur要赋值给pre
			res = append(res, pre)
			pre = cur
		} else { //这个就是重合的情况，重合的情况下要把pre[1]替换pre[1]和re[1]中间最大的那一个
			pre[1] = max(pre[1], cur[1])
		}

	}
	res = append(res, pre)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
