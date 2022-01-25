#incldue <vector>
using namespace std;

vector<vector<int>> generateMatrix(int n) {
	std::vector<std::vector<int>> ret;
	for (int i = 0; i < n; ++i) {
		ret.emplace_back(n);
	}

	int v = 1;
	int q = (n + 1) / 2;
	for (int i = 0; i < q; ++i) {
		int x = i;
		int y = i;
		while (true) {
			ret[x][y] = v++;
			if ((i == q - 1 && n % 2 != 0) || (x == i + 1 && y == i)) {
				break;
			}

			if (x == i && y < n - i - 1) {
				++y;
			}
			else if (y == n - i - 1 && x < n - i - 1) {
				++x;
			}
			else if (x == n - i - 1 && y > i) {
				--y;
			}
			else {
				--x;
			}
		}
	}
	return ret;
}
