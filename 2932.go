func maximumStrongPairXor(nums []int) int {
    abs := func (x int) int {
        if x > 0 {
            return x
        }
        return -x
    }
    ans := 0
    for x := 0; x < len(nums); x++ {
        for y := 0; y < len(nums); y++ {
            if abs(nums[x] - nums[y]) < min(nums[x], nums[y]) {
                ans = max(ans, (nums[x] ^ nums[y]))
            }
        }
    }
    return ans
}
