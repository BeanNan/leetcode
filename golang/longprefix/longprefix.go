package longprefix

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	minLength := computeMinLength(strs)

	var chars []byte
	for i := 0; i < minLength; i++ {
		cur := strs[0][i]
		equalFlag := true
		for _, str := range strs {
			if str[i] != cur {
				equalFlag = false
				break
			}
		}

		if equalFlag {
			chars = append(chars, cur)
		} else {
			break
		}

	}

	return string(chars)
}

func computeMinLength(strs []string) int {

	if len(strs) == 0 {
		return 0
	}

	m := len(strs[0])

	for _, str := range strs {
		if len(str) < m {
			m = len(str)
		}
	}

	return m
}
