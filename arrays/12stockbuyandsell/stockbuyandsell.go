package stockbuyandsell

/*
Return the maximum profit based on the stock's price. 
*/
func StockBuyAndSell(nums []int) int {
	profit := 0

	for i:=1; i<len(nums); i++ {
		if nums[i] > nums[i-1] {
			profit += nums[i] - nums[i-1]
		}
	}

	return profit
}