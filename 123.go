func create_prefix (arr []int) []int {
    dp := make([]int, len(arr) + 1)
    dp[0] = 0
    mn := int(1e7)
    for i := 1; i <= len(arr); i++ { //best that we can do
        dp[i] = max(dp[i - 1], arr[i - 1] - mn)
        mn = min(mn, arr[i - 1])
    }
    return dp
}
func create_suffix (arr []int) []int {
    dp := make([]int, len(arr) + 1)
    dp[len(arr)] = 0
    mx := -int(1e7)
    for i := len(arr) - 1; i >= 0; i-- {
        dp[i] = max(dp[i + 1], mx - arr[i])
        mx = max(mx, arr[i])
    }
    return dp
}
func maxProfit(prices []int) int {
    pref := create_prefix(prices)
    suf := create_suffix(prices)
    ans := pref[len(prices)]
    for i := 0; i < len(prices); i++ {
        ans = max(ans, pref[i] + suf[i])
    }
    return ans
}
