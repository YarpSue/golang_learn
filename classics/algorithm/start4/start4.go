package start4


/*

给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

摩尔投票的方法
1. >n/2 一定是众数
2. 如果记众数的票数+1 非众数的票数-1 则一定所有数字的票数和 >0
3. 若数组的前a个数字的票数和=0 则数组剩下(n-a)个数字的票数和一定仍 >0 即后(n -a) 个数字的众数仍为x


空间复杂度 O(1)
时间复杂度 O(n)
*/

func majorityElement(nums []int) int {
	var votes int
	var d int
	for _, num := range nums {
		if votes == 0 {
			d = num
		}

		if num == d {
			votes++
		} else {
			votes--
		}
	}

	/*
	// 这部分用来判断是不是 大于 ⌊ n/2 ⌋ 的元素   这题是没必要
	var count int
	for _, num := range nums {
		if num == d {
			count++
		}
	}

	if count > len(nums) {
		return d
	} else {
		return 0
	}*/
	return d
}