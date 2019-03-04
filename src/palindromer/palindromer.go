package palindromer

import (

)

func IsPalindrome(s string) bool {
	n := len(s)
	
	for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1 {
		if s[i] != s[j] {
			return false
		}
	}
	
	return true
}