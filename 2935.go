type Trie struct {
    children [2]*Trie
    x int
}
func empty_Trie () *Trie {
    t := Trie{}
    t.init_trie()
    return &t
}
func (root *Trie) init_trie () {
    root.children[0] = nil
    root.children[1] = nil
    root.x = 0
}
func (root *Trie) add_value (x int) {
    root_dummy := root
    for i := 20; i >= 0; i-- {
        root_dummy.x += 1
        b := 0
        if x & (1 << i) != 0 {
            b = 1
        }
        if root_dummy.children[b] == nil {
            root_dummy.children[b] = empty_Trie()
        } 
        root_dummy = root_dummy.children[b]
    }
} 
func (root *Trie) remove_value (x int) {
    root_dummy := root 
    for i := 20; i >= 0; i-- {
        root_dummy.x -= 1
        if root_dummy.x == 0 {
            root_dummy.init_trie()
            break
        }
        b := 0
        if x & (1 << i) != 0 {
            b = 1
        }
        root_dummy = root_dummy.children[b]
    }
}
func (root *Trie) max_xor (x int) int { //max bitwise XOR
    root_dummy := root
    ans := 0
    for i := 20; i >= 0; i-- {
        b := 0
        if x & (1 << i) != 0 {
            b = 1
        }
        if root_dummy.children[1 - b] != nil {
            root_dummy = root_dummy.children[1 - b]
            ans += (1 << i)
        } else if root_dummy.children[b] != nil {
            root_dummy = root_dummy.children[b]
        } else {
            return -1
        }
    }
    return ans
}
func maximumStrongPairXor(nums []int) int {
    t := Trie{}
    t.init_trie()
    ans := 0
    sort.Slice(nums, func(i int, j int) bool {
        return nums[i] < nums[j]
    })
    fmt.Println(nums)
    j := -1
    for x := 0; x < len(nums); x++ {
        for j != len(nums) - 1 && nums[j + 1] <= 2 * nums[x] {
            j += 1
            t.add_value(nums[j])
        }
        ans = max(ans, t.max_xor(nums[x]))
        t.remove_value(nums[x])
    }
    return ans
}
