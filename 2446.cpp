class Solution {
public:
    int time (string s) {
        return (10 * (s[0] - '0') + (s[1] - '0')) * 60 + 10 * (s[3] - '0') + (s[4] - '0');
    }
    bool haveConflict(vector<string>& event1, vector<string>& event2) {
        return !(time(event1[1]) < time(event2[0]) || time(event1[0]) > time(event2[1]));
    }
};
