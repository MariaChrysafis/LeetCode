type SegmentTree struct {
    vec []int64
}
func (st *SegmentTree) build (n int) {
    for n & (n - 1) != 0 {
        n += 1
    }
    st.vec = make([]int64, 2 * n)
}
func (st *SegmentTree) query_rec (dum int, tl int, tr int, l int, r int) int64 {
    if (tl >= l && tr <= r) {
        return st.vec[dum]
    }
    if tl > r || tr < l {
        return -1e15
    }
    return max(st.query_rec(2 * dum + 1, tl, (tl + tr)/2, l, r), st.query_rec(2 * dum + 2, (tl + tr)/2 + 1, tr, l, r))
}
func (st *SegmentTree) query (l int, r int) int64 {
    return st.query_rec(0, 0, len(st.vec)/2 - 1, l, r)
}
func (st *SegmentTree) update (x int, y int64) {
    x += len(st.vec)/2 - 1
    st.vec[x] = y
    for x != 0 {
        x = (x - 1)/2
        st.vec[x] = max(st.vec[2 * x + 1] , st.vec[2 * x + 2])
    }
}
func maxProfit(prices []int, profits []int) int {
    n := len(profits)
    vec := make([][]int, n) 
    for i := 0; i < n; i++ {
        vec[i] = make([]int, 2)
        vec[i][0] = prices[i]
        vec[i][1] = i
    }
    sort.Slice (vec, func(i int, j int) bool {
        if vec[i][0] != vec[j][0] {
            return vec[i][0] < vec[j][0]
        }
        return vec[i][1] < vec[j][1]
    })
    st1 := SegmentTree{} //maintains everyone with smaller price
    st2 := SegmentTree{} //maintains everyone with larger price
    st1.build(n)
    st2.build(n)
    for i := 0; i < n; i++ {
        st1.update(i, int64(-1e9))
        st2.update(i, int64(profits[i]))
    }

    ans := int64(-1e9)
    j1 := 0
    j2 := 0
    fmt.Println(vec)
    for i := 0; i < n; i++ {
        for j1 != n && vec[j1][0] == vec[i][0] {
            st2.update(vec[j1][1], int64(-1e9))
            j1 += 1
        }
        for j2 != n && vec[j2][0] < vec[i][0] {
            st1.update(vec[j2][1], int64(profits[vec[j2][1]]))
            j2 += 1
        }
        ans = max(ans, int64(profits[vec[i][1]]) + st1.query(0, vec[i][1] - 1) + st2.query(vec[i][1] + 1, n - 1))
    }
    if ans < 0 {
        return -1
    }
    return int(ans)
}
