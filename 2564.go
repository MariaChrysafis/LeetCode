func substringXorQueries(s string, queries [][]int) [][]int {
    m := make(map[string]int)
    for i := len(s) - 1; i >= 0; i-- {
        for j := i; j < min(len(s), 30 + i); j++ {
            m[s[i:j + 1]] = i
        }
    }
    ans := make([][]int, len(queries))
    for i, x := range queries {
        s := strconv.FormatInt(int64(x[0] ^ x[1]), 2)
        if val, ok := m[s]; ok {
            ans[i] = []int{val, val + len(s) - 1}
        } else {
            ans[i] = []int{-1, -1}
        }
    }
    return ans
}
