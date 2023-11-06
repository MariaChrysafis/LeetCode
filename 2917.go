func findKOr(nums []int, k int) int {
    ans := 0
    for i := 0; i < 32; i++ {
        sum := 0
        for _, x := range nums {
            if x & (1 << i) == (1 << i) {
                sum += 1
            }
        }
        if sum >= k {
            ans ^= (1 << i)
        }
    }
    return ans
}
