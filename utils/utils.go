package utils

import "strconv"

// General
func Reverse(str string) string {
	if str == "" {
		return str
	}

	lastChar := str[len(str)-1]
	newStr := str[:len(str)-1]

	return string(lastChar) + Reverse(newStr)
}

// Digits in strings
func FindFirstDigit(line string) int {
	for _, char := range line {
		if char >= '0' && char <= '9' {
			digit, err := strconv.Atoi(string(char))
			if err == nil {
				return digit
			}
		}
	}
	return -1
}
