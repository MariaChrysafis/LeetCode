func largestSubmatrix(matrix [][]int) int {
    ans := 0
    res := make([]int, len(matrix[0]))
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            res[j] += 1;
            if matrix[i][j] == 0 {
                res[j] = 0
            } 
        }
        gamma := make([]int, len(res))
        copy(gamma, res)
        sort.Slice(gamma, func(i int, j int) bool {
            return gamma[i] > gamma[j];
        })
        for j := 0; j < len(gamma); j++ {
            ans = max(ans, gamma[j] * (j + 1));
        }
    }
    return ans
}
