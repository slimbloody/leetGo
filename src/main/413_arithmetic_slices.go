package main

type P413 struct {
}

func numberOfArithmeticSlices(nums []int) int {
	var p P413
	return p.sol1(nums)
}

func (p P413)sol1(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	var gaps = make([]int, len(nums) - 1)

	for i := 0; i < len(nums) - 1; i++ {
		gaps[i] = nums[i + 1] - nums[i]
	}
}

func main()  {
}
