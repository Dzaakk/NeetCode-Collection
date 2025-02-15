package main

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := []rune{}

	matchingBrackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != matchingBrackets[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// func main() {
// 	fmt.Println(isValid("()"))      // true
// 	fmt.Println(isValid("()[]{}"))  // true
// 	fmt.Println(isValid("(]"))      // false
// 	fmt.Println(isValid("([)]"))    // false
// 	fmt.Println(isValid("{[]}"))    // true
// 	fmt.Println(isValid(""))        // true (empty string is considered valid)
// 	fmt.Println(isValid("((()))"))  // true
// 	fmt.Println(isValid("{{{{"))    // false
// 	fmt.Println(isValid("}"))       // false
// 	fmt.Println(isValid("([]{})"))  //true
// 	fmt.Println(isValid("([{}])"))  //true
// 	fmt.Println(isValid("([{}])(")) //false
// 	fmt.Println(isValid("([{}])}")) //false
// }
