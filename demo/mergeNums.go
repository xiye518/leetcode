package main

import (
	"fmt"
	"strconv"
)

/*

要求三个及三个以上的数字要合并输出，

输入[1, 2, 3, 5, 6, 8, 9, 10, 13]，
输出["1 -> 3", "5", "6", "8 -> 10", "13"]，

输入[-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20]
输出	"-6,-3-1,3-5,7-11,14,15,17-20"
*/
func convert(nums []int) []string {
	ret := make([]string, 0, len(nums))
	start, end := 0, 0
	var lock bool
	for i := 0; i < len(nums); i++ {
		if i+2 < len(nums) && nums[i]+1 == nums[i+1] && nums[i+1]+1 == nums[i+2] { //三个连续数
			if !lock {
				lock = true
				start = i
			}
			end = i + 2
			continue
		}

		if start != end {
			ret = append(ret, fmt.Sprintf("%d -> %d", nums[start], nums[end]))
			i = end // i的起点位修改
		} else {
			ret = append(ret, strconv.Itoa(nums[i]))
		}
		start, end = i, i
		lock = false
	}

	//fmt.Printf("%#v\n", convert([]int{1, 2, 3, 5, 6, 8, 9, 10, 13}))
	//fmt.Printf("%#v\n", convert([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))
	return ret
}
