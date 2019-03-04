package test_pkg

import (
	"testing"
	"palindromer"
)

func TestPalindrome(t *testing.T) {
	tData := map[string]bool {
		"aabaa": true,
		"aaacc": false,
		"HelleH": true,
		"a": true,
		"ba": false,
		"Hello, olleH": true,
	}
	
	for s, res := range tData {
		if palindromer.IsPalindrome(s) != res {
			t.Errorf("Error: IsPalindrome(%s) == %v", s, palindromer.IsPalindrome(s))
		}
	}
}
