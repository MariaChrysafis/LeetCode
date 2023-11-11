func distributeCandies(n int, limit int) int {
    ans := 0
    for i := 0; i <= limit; i++ {
        right := min(limit, n - i)
        left := max(0, n - i - limit)
        if left > right {
            continue
        }
        ans += right - left + 1
    }
    return ans
}
