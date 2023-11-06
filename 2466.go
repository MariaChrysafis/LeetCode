func countGoodStrings(low int, high int, zero int, one int) int {
    sub := func (a int, b int) int {
        return (a - b + 1e9 + 7) % (1e9 + 7)
    }
    mult := func (a int, b int) int {
        return (a * b) % (1e9 + 7)
    }
    add := func (a int, b int) int {
        if a + b >= 1e9 + 7 {
            return a + b - 1e9 - 7
        } else {
            return a + b
        }
    }
    if zero < one {
        zero, one = one, zero
    }
    fact := make([]int, high + 10)
    for i := 0; i < len(fact); i++ {
        if i == 0 {
            fact[i] = 1
        } else {
            fact[i] = mult(i, fact[i - 1])
        }
    }
    binPow := func (a int, y int) int {
        ans := 1
        res := a
        for y != 0 {
            if y % 2 == 1 {
                ans = mult(ans, res)
            }
            res = mult(res, res)
            y /= 2
        }
        return ans
    }
    inv := func(a int) int {
        return binPow(a, 1e9 + 7 - 2)
    }
    combo := func(a int, b int) int {
        return mult(fact[a], inv(mult(fact[b], fact[a - b])))
    }
    cntr := 0
    for z := 0; z * zero <= high; z++ {
        x := (low - z * zero + one - 1)/one
        y := (high - z * zero)/one
        if x > y {
            continue
        }
        if x >= 1 {
            cntr = sub(cntr, combo(x + z, z + 1))
        }
        cntr = add(cntr, combo(y + 1 + z, z + 1))
    }
    return cntr
}
