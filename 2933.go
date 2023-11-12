func to_int (s string) int {
    hour := 10 * (int(s[0]) - int('0')) + int(s[1]) - int('0')
    minute := 10 * (int(s[2]) - int('0')) + int(s[3]) - int('0')
    return 60 * hour + minute
}
func findHighAccessEmployees(access_times [][]string) []string {
    m := make(map[string][]string)
    for _, x := range access_times {
        m[x[0]] = append(m[x[0]], x[1])
    }
    ans := make([]string, 0)
    for key, value := range m {
        sort.Slice (value, func(i int, j int) bool {
            return value[i] < value[j]
        })
        for i := 0; i + 2 < len(value); i++ {
            if to_int(value[i + 2]) - to_int(value[i]) < 60 {
                ans = append(ans, key)
                break
            }
        }
    }
    return ans
}
