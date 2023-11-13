func minimumScore(s string, t string) int {
    pref := make([]int, len(t) + 1) 
    suf := make([]int, len(t) + 1)
    for i := len(suf) - 1; i >= 0; i-- {
        if i == len(suf) - 1 {
            suf[i] = len(s)
        } else {
            x := suf[i + 1] - 1
            for x >= 0 && s[x] != t[i] {
                x--
            }
            if x < 0 {
                suf[i] = -1
            } else {
                suf[i] = x
            }
        }
    }
    for i := 0; i < len(pref); i++ {
        if i == 0 {
            pref[i] = -1 
        } else {
            if pref[i - 1] == len(s) {
                pref[i] = len(s)
                continue
            }
            x := pref[i - 1] + 1
            for x != len(s) && s[x] != t[i - 1] {
                x += 1
            }
            pref[i] = x
        }
    }
    ans := len(t)
    for i := 0; i < len(pref); i++ {
        if pref[i] >= suf[len(pref) - 1] {
            continue
        }
        l := i
        r := len(pref) - 1
        for l != r {
            m := (l + r)/2
            if suf[m] > pref[i] {
                r = m
            } else {
                l = m + 1
            }
        }
        if pref[i] < suf[l] {
            ans = min(ans, l - i)
        }
    }
    return ans
}
