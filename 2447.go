func gcd (a int, b int) int {
    if a == 0 || b == 0 {
        return max(a, b);
    }
    return gcd(max(a, b) % min(a, b), min(a, b));
}
func subarrayGCD(nums []int, k int) int {
    ans := 0
    for i := 0; i < len(nums); i++ {
        g := 0
        for j := i; j < len(nums); j++ {
            g = gcd(g, nums[j])
            if g == k {
                ans += 1
            }
        }
    }
    return ans
}
