package longest_palindromic_substring

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	res := string(s[0])
	for i, _ := range s {
		if tmp, valid := isPalindrome(&s, i, i); valid {
			if len(tmp) > len(res) {
				res = tmp
			}
		}
		if tmp, valid := isPalindrome(&s, i, i+1); valid {
			if len(tmp) > len(res) {
				res = tmp
			}
		}
	}

	return res
}

func isPalindrome(s *string, l int, r int) (string, bool) {
	for l >= 0 && r < len(*s) && (*s)[l] == (*s)[r] {
		l--
		r++
	}
	l++
	r--
	valid := r > l
	if !valid {
		return "", valid
	}

	return string((*s)[l : r+1]), valid
}
