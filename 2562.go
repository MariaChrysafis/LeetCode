func findTheArrayConcVal(nums []int) int64 {
    ans := int64(0)
    for i := 0; i < len(nums)/2; i++ {
        x, _ := strconv.ParseInt(strconv.Itoa(nums[i]) + strconv.Itoa(nums[len(nums) - 1 - i]), 10, 64)
        ans += x
    }
    if len(nums) % 2 == 1 {
        ans += int64(nums[len(nums)/2])
    }
    return ans
}
