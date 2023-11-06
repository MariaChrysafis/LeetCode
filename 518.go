func change(amount int, coins []int) int {
    dp := make([][]int, amount + 1)
    for i := 0; i <= amount; i++ {
        dp[i] = make([]int, len(coins) + 1)
    }
    for i := 0; i <= amount; i++ {
        for j := 0; j < len(coins); j++ {
            if i == 0 {
                dp[i][j] = 1
            } else {
                dp[i][j] = 0
                if j != 0 {
                    dp[i][j] = dp[i][j - 1]
                }
                if i >= coins[j] {
                    dp[i][j] += dp[i - coins[j]][j]
                }
            }
        }
    }
    return dp[amount][len(coins) - 1]
}
