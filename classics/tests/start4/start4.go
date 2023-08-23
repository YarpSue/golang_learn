package start4

/*
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。


*/


func majorityElement(nums []int) int {

	sum := 0
	p := 0
	for _, n := range nums {
		if sum == 0 {
			p = n
		}

		if n == p {
			sum++
		} else {
			sum--
		}

	}
	return p
}