struct SparseTable {
    vector<vector<int>> dp;
    int query (int l, int r) {
        int sz = log2(r - l + 1);
        return max(dp[l][sz], dp[r - (1 << sz) + 1][sz]);
    }
    SparseTable (vector<int> v) {
        int n = v.size();
        dp.resize(v.size());
        for (int i = 0; i < v.size(); i++) {
            dp[i].resize(30);
            dp[i][0] = v[i];
        }
        for (int j = 1; j < dp[0].size(); j++) {
            for (int i = 0; i < v.size(); i++) {
                dp[i][j] = max(dp[i][j - 1], dp[min(i + (1 << (j - 1)), n - 1)][j - 1]);
            }
        }
    }
};
class Solution {
public:
    vector<int> leftmostBuildingQueries(vector<int>& heights, vector<vector<int>>& queries) {
        SparseTable st(heights);
        vector<int> v;
        for (auto& q : queries) {
            int alice = q[0];
            int bob = q[1];
            int res = INT_MAX;
            if (alice > bob) {
                swap(alice, bob);
            }
            if (alice == bob) {
                v.push_back(alice);
                continue;
            }
            if (alice <= bob and heights[bob] > heights[alice]) {
                res = bob;
            }
            int m = max(heights[alice], heights[bob]);
            if (st.query(max(alice, bob), heights.size() - 1) <= m) {
                if (res == INT_MAX) {
                    v.push_back(-1);
                } else {
                    v.push_back(res);
                }
                continue;
            }
            int left = max(alice, bob);
            int right = heights.size() - 1;
            while (left != right) {
                int mid = (left + right)/2;
                if (st.query(max(alice, bob), mid) > m) {
                    right = mid;
                } else {
                    left = mid + 1;
                }
            }
            res = min(res, left);
            v.push_back(res);
        }
        return v;
    }
};
