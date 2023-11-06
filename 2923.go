func findChampion(grid [][]int) int {
    for i := 0; i < len(grid); i++ {
        sum := 0
        for j := 0; j < len(grid); j++ {
            sum += grid[i][j]
        }
        if sum == len(grid) - 1 {
            return i
        }
    }
    return -1
}
