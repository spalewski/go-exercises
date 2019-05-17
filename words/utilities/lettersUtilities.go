package utilities

import (
	"fmt"
	"strings"
)

func IsVovel(x string) bool {
	vowels := [9]string{"a", "e", "i", "o", "ą", "y", "ę", "ó", "u"}
	vowelLookupTable := make(map[string]bool)
	for _, v := range vowels {
		vowelLookupTable[v] = true
	}
	return vowelLookupTable[x]
}

func CountVovels(input string) int {
	vovelCount := 0
	var lowerCase string = strings.ToLower(input)
	for _, x := range lowerCase {
		var s string
		s = fmt.Sprintf("%c", x)
		if IsVovel(s) {
			vovelCount += 1
		}
	}
	return vovelCount
}

func IsEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}
