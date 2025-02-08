package main

// import (
// 	"fmt"
// )

func isPalindrome(s string) bool {
	newStr := ""

	for _, c := range s {
		if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
			newStr += string(c)
		} else if 'A' <= c && c <= 'Z' {
			newStr += string(c + 'a' - 'A')
		}
	}

	reversedStr := reverse(newStr)
	return newStr == reversedStr
}

func reverse(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

// func main() {
// 	str1 := "Was it a car or a cat I saw?" //true
// 	str2 := "tab a cat"                    //false
// 	fmt.Println(isPalindrome(str1))
// 	fmt.Println(isPalindrome(str2))
// }
