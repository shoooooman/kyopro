#include <iostream>
#include <vector>

using namespace std;

int main() {
    int n, m;
    cin >> n >> m;
    vector<double> a(n);
    double sum = 0.0;
    for (int i = 0; i < n; i++) {
        cin >> a[i];
        sum += a[i];
    }
    int count = 0;
    for (double v : a) {
        if (v >= sum/(4.0*m)) {
            count++;
            if (count >= m) {
                cout << "Yes" << endl;
                return 0;
            }
        }
    }
    cout << "No" << endl;
}
