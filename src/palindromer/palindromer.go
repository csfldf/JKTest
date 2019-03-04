package palindromer

import (
//	"unicode"
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

//func IsPalindrome(s string) bool {
//	n := len(s)
//	
//	i, j := 0, n - 1
//	
//	for i < j {
//		for i < j && !unicode.IsDigit(rune(s[i])) && !unicode.IsLetter(rune(s[i])) {
//			i++
//		}
//		
//		for i < j && !unicode.IsDigit(rune(s[j])) && !unicode.IsLetter(rune(s[j])) {
//			j--
//		}
//		
//		if i < j && s[i] != s[j] {
//			return false
//		} else {
//			i++
//			j--
//		}
//	}
//	
//	return true
//}