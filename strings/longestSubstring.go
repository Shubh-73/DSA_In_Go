package strings

func LongestPalindromicSubstring(s string) string {

	if len(s) == 0 {
		return ""
	}

	start, maxLength := 0, 0

	expandaAroundCenter := func(left, right int) {

		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++

		}

		if right-left-1 > maxLength {
			maxLength = right - left
			start = left + 1
		}

	}

	for i := 0; i < len(s); i++ {
		expandaAroundCenter(i, i)
		expandaAroundCenter(i, i+1)
	}

	return s[start : start+maxLength]
}
