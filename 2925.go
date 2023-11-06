func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
    //dp[x] -> answer for the subtree of x
    n := len(values)
    dp := make([]int64, n)
    sub := make([]int64, n)
    adj := make([][]int, n)
    for i:=0; i < n; i++ {
        adj[i] = make([]int, 0)
    }
    for _, x := range edges {
        adj[x[0]] = append(adj[x[0]], x[1])
        adj[x[1]] = append(adj[x[1]], x[0])
    }
    //fill up dp
    var dfs func(curNode int, prevNode int)
    dfs = func (curNode int, prevNode int) {
        if len(adj[curNode]) == 1 && curNode != prevNode {
            dp[curNode] = 0
            sub[curNode] = int64(values[curNode])
            return
        }
        dp[curNode] = int64(values[curNode])
        sub[curNode] = int64(values[curNode])
        for _, i := range adj[curNode] {
            if i != prevNode {
                dfs (i, curNode)
                dp[curNode] += dp[i]
                sub[curNode] += sub[i]
            }
        }
        dp[curNode] = max(dp[curNode], sub[curNode] - int64(values[curNode]))
    }
    dfs(0, 0)
    return dp[0]
}
