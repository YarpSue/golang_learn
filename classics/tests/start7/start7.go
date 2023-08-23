package start7

/*

给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。

*/


func canJump(nums []int) bool {

	rightMost := 0
	numLen := len(nums)

	for index, num := range nums {

		if index <= rightMost {
			if index + num > rightMost {
				rightMost = index + num
			}
			if rightMost >= numLen-1 {
				return true
			}
		}
	}

	return false
}