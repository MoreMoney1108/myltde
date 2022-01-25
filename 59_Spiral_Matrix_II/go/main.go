package spiralmmatrix2

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}

	v := 1
	q := (n + 1) / 2
	for i := 0; i < q; i++ {
		x := i
		y := i

		for {
			ret[x][y] = v
			v++
			if (i == q-1 && n%2 != 0) || (x == i+1 && y == i) {
				break
			}

			switch {
			case x == i && y < n-i-1:
				y++
			case y == n-i-1 && x < n-i-1:
				x++
			case x == n-i-1 && y > i:
				y--
			default:
				x--
			}
		}
	}

	return ret
}
