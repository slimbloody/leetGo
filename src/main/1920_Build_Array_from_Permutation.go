package main

import "fmt"

type P1920 struct {
}

/*
 */

func buildArray(nums []int) []int {
	var p P1920
	return p.sol1(nums)
}

/*
elite:
in-place algorithm

https://leetcode.com/problems/build-array-from-permutation/discuss/1315926/Python-O(n)-Time-O(1)-Space-w-Full-Explanation

keypoint: 0 <= i < nums.length

b = a // q (where // is integer division) - we know that qb when divided by q will give us b, however we still would need to get rid of the r // q. From our requirements though, r < q, so r // q will always be 0, thus b = (qb//q) + (r//q) = b + 0 = b

r = a % q - we know that qb is a multiple of q, thus is divided by it cleanly and we know that r < q, so r is not a multiple of q, therefore the remainder when dividing a = qb + r by q is just r

 */

func (p P1920)sol1(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		nums[i] = nums[i] + n * (nums[nums[i]] % n)
	}

	for i := 0; i < n; i++ {
		nums[i] =nums[i] / n
	}

	return nums
}

func (p P1920)sol2(nums []int) []int {
	var res []int
	for _, num := range nums {
		res = append(res, nums[num])
	}

	return res
}

func main()  {
	// [0,1,2,4,5,3]
	fmt.Println(buildArray([]int{0,2,1,5,3,4}))

	// [4,5,0,1,2,3]
	fmt.Println(buildArray([]int{5,0,1,2,3,4}))
}