class Solution {
public:
    long long maxSpending(vector<vector<int>>& values) {
        set<pair<int,pair<int,int>>> mySet; //(value, index)
        for (int i = 0; i < values.size(); i++) {
            mySet.insert(make_pair(values[i][values[0].size() -1], make_pair(i, values[0].size() - 1)));
        }
        int64_t ans = 0;
        int cntr = 0;
        while (!mySet.empty()) {
            pair<int,pair<int,int>> p = *mySet.begin();
            ans += (int64_t)p.first * (int64_t)(++cntr);
            mySet.erase(p);
            if (p.second.second == 0) {
                continue;
            }
            mySet.insert(make_pair(values[p.second.first][p.second.second - 1], make_pair(p.second.first, p.second.second - 1)));
        }
        return ans;
    }
};
