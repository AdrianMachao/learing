package leetcode

func ztrans(s string, n int) string {
	count := 2 + 2*(n-2)
	ans := make([]byte, len(s))
	for i := 0; i <= count/2; i++ {
		if i == 0 || i == count/2 {
			for j := i; i < len(s); j = j + count {
				ans = append(ans, s[j])
			}
			continue
		}
		// if i== count/2 {
		// 	for j := i; i < len(s); j = j + count {
		// 		ans = append(ans, s[j])
		// 	}
		// }
		// for m=i;i
		for j := 0; j < len(s); j = j + count {
			if j+i < len(s) {
				ans = append(ans, s[j])
			}
			if j+count-i < len(s) {
				ans = append(ans, s[j])
			}
		}
	}

	return string(ans)
}
