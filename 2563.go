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
        return 0
    }
    return st.query_rec(2 * dum + 1, tl, (tl + tr)/2, l, r) + st.query_rec(2 * dum + 2, (tl + tr)/2 + 1, tr, l, r)
}
func (st *SegmentTree) query (l int, r int) int64 {
    return st.query_rec(0, 0, len(st.vec)/2 - 1, l, r)
}
func (st *SegmentTree) update (x int, y int64) {
    x += len(st.vec)/2 - 1
    st.vec[x] += y
    for x != 0 {
        x = (x - 1)/2
        st.vec[x] = st.vec[2 * x + 1] + st.vec[2 * x + 2]
    }
}
func countFairPairs(nums []int, lower int, upper int) int64 {
    new_nums := make([]int, 0)
    new_nums = append(new_nums, 0)
    for _, x := range nums {
        new_nums = append(new_nums, x)
        new_nums = append(new_nums, upper - x) //the bound here
        new_nums = append(new_nums, lower - x)
    }
    sort.Slice(new_nums, func(i int, j int) bool {
        return new_nums[i] < new_nums[j]
    })
    myMap := make(map[int]int)
    cntr := 0
    for i := 0; i < len(new_nums); i++ {
        myMap[new_nums[i]] = cntr
        cntr += 1
    }
    //create a new segment tree 
    st := SegmentTree{}
    st.build(len(new_nums))
    ans := int64(0)
    for i := 0; i < len(nums); i++ {
        ans += st.query(myMap[lower - nums[i]], myMap[upper - nums[i]])
        st.update(myMap[nums[i]], 1)
    }
    return ans
}
