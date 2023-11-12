func min1 (x int, y int) int {
    if min(x, y) == -1 {
        return max(x, y)
    }
    return min(x, y)
}
func trier (nums1 []int, nums2[] int) int {
    ans := 0
    n := len(nums1)
    if nums1[n - 1] > nums2[n - 1] {
        nums1[n - 1], nums2[n - 1] = nums2[n - 1], nums1[n - 1]
        ans += 1
    }
    for i := 0; i < n; i++ {
        if nums1[i] > nums1[n - 1] || nums2[i] > nums2[n - 1] {
            nums1[i], nums2[i] = nums2[i], nums1[i]
            ans += 1
        }
        if nums1[i] > nums1[n - 1] || nums2[i] > nums2[n - 1]{
            return -1
        }
    }
    return ans
}
func minOperations(nums1 []int, nums2 []int) int {
    n1 := make([]int, len(nums1))
    n2 := make([]int, len(nums2))
    copy(n1, nums1)
    copy(n2, nums2)
    return min1(trier(nums1, nums2), trier(n2, n1))
}
