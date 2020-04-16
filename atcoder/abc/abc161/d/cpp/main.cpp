#include <iostream>
#include <vector>
#include <string>

using namespace std;

vector<string> getDiff1(string str)  {
    vector<string> diff1(0);
    int last = str[str.length()-1] - '0';
    if (last != 0) {
        diff1.push_back(to_string(last - 1));
    }
    diff1.push_back(to_string(last));
    if (last != 9) {
        diff1.push_back(to_string(last + 1));
    }
    return diff1;
}

int main() {
    int k;
    cin >> k;

    long long count = 0;
    int digit = 1;
    vector<vector<string> > lunlun(1, vector<string>(0));
    for (int i = 1; i <= 9; i++) {
        lunlun[0].push_back(to_string(i));
        count++;
        if (count == k) {
            cout << i << endl;
            return 0;
        }
    }

    while(1) {
        lunlun.push_back(vector<string>(0));
        for (int i = 0; i < lunlun[digit-1].size(); i++) {
            vector<string> diff1 = getDiff1(lunlun[digit-1][i]);
            for (string diff : diff1) {
                lunlun[digit].push_back(lunlun[digit-1][i] + diff);
                count++;
                if (count == k) {
                    int last_index = lunlun[digit].size() - 1;
                    long long ans = stoll(lunlun[digit][last_index]);
                    cout << ans << endl;
                    return 0;
                }
            }
        }
        digit++;
    }
    return 0;
}
