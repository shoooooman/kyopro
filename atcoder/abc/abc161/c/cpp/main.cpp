#include <iostream>

using namespace std;

int main() {
    long long n, k;
    cin >> n >> k;

    long long mod = n % k;
    long long mod2 = k - mod;
    cout << min(mod, mod2) << endl;
}
