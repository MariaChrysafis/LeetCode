func distributeCandies(n int, limit int) int64 {
    ans := int64(0)
    l1 := int64(max(0, n - 2 * limit))
    r1 := int64(min(n - limit, limit))
    if l1 <= r1 {
        ans += int64(2 * limit + 1 - n) * int64(r1 - l1 + 1)
        ans += r1 * (r1 + 1)/2 - l1 * (l1 - 1)/2
    }
    l1 = int64(max(n - limit + 1, 0))
    r1 = int64(min(limit, n))
    if l1 <= r1 {
        ans += int64(n + 1) * (r1 - l1 + 1)
        ans -= r1 * (r1 + 1)/2 - l1 * (l1 - 1)/2
    }
    return ans
}
