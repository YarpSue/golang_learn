package main

import (
	"fmt"
	"math"
)

/*
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。


*/

func maxProfit(prices []int) int {
	cost := math.MaxInt64
	profit := 0
	for _, price := range prices {
		if price < cost {
			cost = price
		}
		if (price - cost) > profit {
			profit = price - cost
		}
	}

	return profit
}

func main() {
	//x := []int{9,3,12,1,2,3}
	x := []int{7,1,5,3,6,4}
	y := maxProfit(x)
	fmt.Println(y)
}


/*
给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。

在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。

返回 你能获得的 最大 利润 。

贪心算法能过
*/


func maxProfit2(prices []int) int {
	cost := math.MaxInt64
	profit := 0
	for _, price := range prices {
		if price < cost {
			cost = price
		}
		if (price - cost) > 0 {
			profit += price - cost
			cost = price
		}
	}

	return profit
}