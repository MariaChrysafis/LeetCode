func goodBinaryStrings(minLength int, maxLength int, oneGroup int, zeroGroup int) int {
    dp := make([]int, maxLength + 1)
    for i := 0; i < len(dp); i += 1 {
        if i == 0 {
            dp[i] = 1
        } else {
            dp[i] = 0
            if i >= zeroGroup {
                dp[i] += dp[i - zeroGroup]
            }
            if i >= oneGroup {
                dp[i] += dp[i - oneGroup]
            }
            dp[i] %= (1e9 + 7)
        }
    }
    ans := 0
    for i := minLength; i <= maxLength; i++ {
        ans += dp[i]
        ans %= (1e9 + 7)
    }
    return ans
}
