package start6


/*
给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

0 <= j <= nums[i]
i + j < n
返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。

*/


func jump(nums []int) int {

	var maxPos int  // 现在能跳到最大的位置
	var step int
	var maxPosPointer int // 前一个maxPosPointer
	for i, num := range nums {

		if i == len(nums) -1 {
			break
		}

		maxPos = max(maxPos, i+num)
		if i == maxPosPointer {  // 循环指针移动到 maxPosPointer 说明是 能本次能跳到的最远位置  然后更新
			// 如果最远跳到的是最后一个元素  那么遍历就会多跳一次  步数会多加一次  没有必要  因为最后一个元素就是本次跳的目标
			// 所以加上前面那个 break程序
			maxPosPointer = maxPos
			step++
		}
	}
	return step
}

func max(a int, b int) int{

	if b > a {
		return b
	}

	return a
}