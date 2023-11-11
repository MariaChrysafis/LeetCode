func stringCount(n int) int {
    dp := make([][][][]int64, n + 1)
    for i := 0; i <= n; i++ {
        dp[i] = make([][][]int64, 3)
        for e := 0; e <= 2; e++ {
            dp[i][e] = make([][]int64, 2)
            for t := 0; t <= 1; t++ {
                dp[i][e][t] = make([]int64, 2)
                for c := 0; c <= 1; c++ {
                    if i == 0 && e == 0 && t == 0 && c == 0 {
                        dp[i][e][t][c] = 1
                    } else if i == 0 {
                        dp[i][e][t][c] = 0 
                    } else {
                        dp[i][e][t][c] += dp[i - 1][e][t][c] * 23
                        if e != 0 {
                            dp[i][e][t][c] += dp[i - 1][e - 1][t][c]
                        }
                        if e == 2 {
                            dp[i][e][t][c] += dp[i - 1][e][t][c]
                        }
                        if t != 0 {
                            dp[i][e][t][c] += dp[i - 1][e][t - 1][c]
                        }
                        if t == 1 {
                            dp[i][e][t][c] += dp[i - 1][e][t][c]
                        }
                        if c != 0 {
                            dp[i][e][t][c] += dp[i - 1][e][t][c - 1]
                        }
                        if c == 1 {
                            dp[i][e][t][c] += dp[i - 1][e][t][c]
                        }
                        dp[i][e][t][c] %= int64(1e9 + 7)
                    }
                }
            }
        }
    }
    return int(dp[n][2][1][1])
}
