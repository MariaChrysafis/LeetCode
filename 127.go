type Graph struct {
    adj [][]int
}
type Pair struct {
    node int
    distance int
}
func (gr *Graph) build (n int) {
    gr.adj = make([][]int, n)
    for i := 0; i < n; i++ {
        gr.adj[i] = make([]int, 0)
    }
}
func (gr *Graph) add_edge (u int, v int) {
    gr.adj[u] = append(gr.adj[u], v)
    gr.adj[v] = append(gr.adj[v], u)
}
func similar (s1 string, s2 string) bool {
    ind := 0
    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i] {
            ind++
        }
    }
    return (ind == 1)
}
func (gr *Graph) shortestPath (start int, end int) int {
    q := make([]Pair, 0)
    q = append(q, Pair{start, 0})
    dist := make([]int, len(gr.adj))
    for i := 0; i < len(dist); i++ {
        dist[i] = -1
    }
    dist[start] = 0
    for len(q) != 0 {
        x := q[0]
        q = q[1:]
        for _, new_node := range gr.adj[x.node] {
            if dist[new_node] == -1 {
                dist[new_node] = x.distance + 1
                q = append(q, Pair{new_node, dist[new_node]})
            }
        }
    }
    return dist[end]
}
func contains (arr []string, x string) bool {
    for _, i := range arr {
        if i == x {
            return true
        }
    } 
    return false
}
func ladderLength(beginWord string, endWord string, wordList []string) int {
    if !contains(wordList, endWord) {
        return 0
    }
    wordList = append(wordList, beginWord)
    wordList = append(wordList, endWord)
    gr := Graph{}
    gr.build(len(wordList))
    for i := 0; i < len(wordList); i++ {
        for j := 0; j < i; j++ {
            if similar(wordList[i], wordList[j]) {
                gr.add_edge(i, j)
            }
        }
    }
    return gr.shortestPath(len(wordList) - 1, len(wordList) - 2) + 1
}
