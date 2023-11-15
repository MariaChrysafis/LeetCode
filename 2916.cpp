template<class T>
class SegmentTree {
public:

    SegmentTree (int N) {
        N = (1 << ((int)floor(log2(N - 1)) + 1));
        this->N = N;
        val.assign(2 * N, ID);
    }

    void update (int x, T y) {
        x += N - 1;
        val[x] = y;
        while (x != 0) {
            x = (x - 1)/2;
            val[x] = merge(val[2 * x + 1], val[2 * x + 2]);
        }
    }

    T query (int ind, const int l, const int r, int tl, int tr) {
        if (tl >= l && tr <= r) {
            return val[ind];
        }
        if (tr < l || tl > r) {
            return ID;
        }
        return merge(query(2 * ind + 1, l, r, tl, (tl + tr)/2), query(2 * ind + 2, l, r, (tl + tr)/2 + 1, tr));
    }

    T query (int l, int r) {
        return query(0, l, r, 0, N - 1);
    }
private:
    vector<T> val;
    T ID = 0;
    T merge (T x, T y) {
        return x + y;
    }
    int N;
};
class Solution {
public:
    int sumCounts(vector<int>& nums) {
        const int MOD = 1e9 + 7;
        int64_t n = nums.size();
        map<int,vector<int>> oc;
        for (int i = 0; i < n; i++) {
            oc[nums[i]].push_back(i);
        }
        for (auto& p: oc) {
            reverse(p.second.begin(), p.second.end());
        }
        int64_t sum = 0;
        /*
        auto get_cost = [&] () {
            return n * x.size() * x.size() + sum;
            int64_t ind = 0;
            for (int64_t i: x) {
                ind++;
                ans -= 2 * i * ind;
            }
            return ans;
        };
        */
        SegmentTree<int64_t> st(n + 1);
        SegmentTree<int> st1(n + 1);
        for (int i = 0; i <= n; i++) {
            st.update(i, 0);
            st1.update(i, 0);
        }
        int64_t res = 0;
        auto add = [&] (int64_t val) {
            sum += val;
            st.update(val, val);
            st1.update(val, 1);
            res -= st.query(val + 1, n) * 2;
            res += 2 * MOD, res %= MOD;
            res -= ((val * st1.query(0, val)) % MOD * 2) % MOD;
            res += MOD, res %= MOD;
        };
        auto remove = [&] (int64_t val) {
            sum -= val;
            st.update(val, 0);
            st1.update(val, 0);
            res += st.query(val, n) * 2;
            res += val * (st1.query(0, val) + 1) * 2;
        };
        int64_t ans = 0;
        for (int i = 0; i < n; i++) {
            if (i != 0) {
                remove(i - 1);
                oc[nums[i - 1]].pop_back();
                if (!oc[nums[i - 1]].empty()) {
                    add(oc[nums[i - 1]].back());
                }
            } else {
                set<int> s;
                for (int j = 0; j < nums.size(); j++) {
                    if (not s.count(nums[j])) {
                        s.insert(nums[j]);
                        add(j);
                    }
                }
            }
            ans += ((n * st1.query(0, n)) % MOD * st1.query(0, n)) % MOD + sum % MOD;
            ans += res;
            ans %= MOD;
        }
        
        return ans;
    }
};
