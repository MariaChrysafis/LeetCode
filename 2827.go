func numberOfBeautifulIntegers(low int, high int, k int) int {
    myMap := [11][20][21][2]map[int]int{}
    for i := 0; i < len(myMap); i++ {
        for j := 0; j < len(myMap[i]); j++ {
            for l := 0; l < len(myMap[i][l]); l++ {
                myMap[i][j][l][0] = make(map[int]int)
                myMap[i][j][l][1] = make(map[int]int)
            }
        }
    }
    powr := make([]int, 10)
    powr[0] = 1
    for i := 1; i <= 9; i++ {
        powr[i] = powr[i - 1] * 10
    }
    var countNum func(dig int, high int, eo int, mod int, lead int) int
    countNum = func (dig int, high int, eo int, mod int, lead int) int {
        if high == 0 && dig == -1 {
            if mod == 0 && eo == 0 {
                return 1
            }
            return 0
        }
        if val, ok := myMap[dig + 1][eo + 10][mod][int(lead)][high]; ok {
            return val
        }
        ans := 0
        powr := powr[dig]
        for i := 0; powr * i <= high && i <= 9; i++ {
            new_eo := eo
            new_mod := (mod + powr * i) % k
            new_lead := lead
            new_powr := powr - 1
            if high/powr == i {
                new_powr = high - powr * i
            }
            if i != 0 {
                new_lead = 0
            }
            if i % 2 == 1 {
                new_eo -= 1
            } else if !(lead == 1 && i == 0){
                new_eo += 1
            }
            ans += countNum(dig - 1, new_powr, new_eo, new_mod, new_lead)
        }
        myMap[dig + 1][eo + 10][mod][int(lead)][high] = ans
        return ans
    }
    return countNum(9, high, 0, 0, 1) - countNum(9, low - 1, 0, 0, 1)
}
