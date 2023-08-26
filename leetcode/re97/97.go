package l97

// re: memoではやくした
// i1 * i2までしか処理しないので、O(s1*s2)

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) == 0 {
		return true
	}
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	visited := make([]bool, len(s1)+1)
	return canObtain(&strs{
		s1: s1,
		s2: s2,
		s3: s3,
	},
		len(s1)-1,
		len(s2)-1,
		len(s3)-1,
		visited,
	)
}

type strs struct {
	s1, s2, s3 string
}

func canObtain(strs *strs, i1, i2, i3 int, visited []bool) bool {
	if visited[i1+1] {
		return false
	}
	visited[i1+1] = true

	if i3 == 0 && i1 == 0 && i2 == -1 {
		return strs.s3[0] == strs.s1[0]
	}
	if i3 == 0 && i1 == -1 && i2 == 0 {
		return strs.s3[0] == strs.s2[0]
	}

	return (i1 >= 0 &&
		strs.s3[i3] == strs.s1[i1] &&
		canObtain(strs, i1-1, i2, i3-1, visited)) ||
		(i2 >= 0 &&
			strs.s3[i3] == strs.s2[i2] &&
			canObtain(strs, i1, i2-1, i3-1, visited))
}
