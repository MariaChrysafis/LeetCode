func maxDistance(arrays [][]int) int {
    new_arr := make([][2]int, len(arrays))
    for i, x := range arrays {
        sort.Slice(x, func(i int, j int) bool {
            return x[i] < x[j]
        })
        new_arr[i] = [2]int{x[0], x[len(x) - 1]}
    }
    sort.Slice(new_arr, func(i int, j int) bool {
        return new_arr[i][1] < new_arr[j][1]
    })
    ans := 0
    for i := 0; i < len(new_arr); i++ {
        mx := new_arr[len(new_arr) - 1][1]
        if i == len(new_arr) - 1 {
            mx = new_arr[len(new_arr) - 2][1]
        }
        ans = max(ans, mx - new_arr[i][0])
    }
    return ans
}
