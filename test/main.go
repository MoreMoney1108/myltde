package test

/*
"www.bytedance.com"
"com.bytedance.www"
["w", "w", "w", ".", "b", "y", "t", "e", "d", "a", "n", "c", "e", ".", "c", "o", "m"]
["c", "o", "m", ".", "b", "y", "t", "e", "d", "a", "n", "c", "e", ".", "w", "w", "w"]
*/

func t1(src []string) {
	l := 0
	r := 0
	for r = len(src) - 1; r >= 0; r-- {
		if src[r] == "." {
			r++
			break
		}
	}

	for src[l] != "." {
		src[l], src[r] = src[r], src[l]
		l++
		r++
	}
}
