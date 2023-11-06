func findChampion(n int, edges [][]int) int {
    adj := make([][]int, n)
    for i := 0; i < n; i++ {
        adj[i] = make([]int, 0)
    }
    for _, x := range edges {
        adj[x[0]] = append(adj[x[0]], x[1])
    }
    vis := make([]bool, n)
    var dfs func(int) 
    dfs = func (x int) {
        if vis[x] {
            return;
        }
        vis[x] = true
        for _, i := range adj[x] {
            dfs(i)
        }
    }
    refill := func () {
        for i := 0; i < n; i++ {
            vis[i] = false
        }
    }
    for i := 0; i < n; i++ {
        refill()
        dfs(i)
        fine := true
        for j := 0; j < n; j++ {
            if !vis[j] {
                fine = false
            }
        }
        if fine {
            return i
        }
    }
    return -1
}
