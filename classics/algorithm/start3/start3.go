package main

import "fmt"

/*
给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：

更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
返回 k 。


时间 O(n)
空间 O(1)

*/

func removeDuplicates(nums []int) int {

	p := 0
	for q, num := range nums {
		if q == 0 {
			continue
		}

		if nums[p] != nums[q] {
			if q - p > 1 {   // 优化的点
				nums[p+1] = num
			}
			p++
		}
	}

	return p+1
}



/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

这是另一道题目

时间 O(n)
空间 O(1)
*/


func removeDuplicates2(nums []int) int {

	p := 0
	if len(nums) <= 2 {
		return len(nums)
	}

	for q, num := range nums {
		if q == 0 || q == 1 {
			continue
		}

		if nums[p] != nums[q] {
			if q - p > 2 {
				nums[p+2] = num
			}
			p++
		}
	}

	return p+2
}

func main() {
	s := []int{0,1,1,1,2,2,3,3,3}
	x := removeDuplicates2(s)
	fmt.Println(x)
}