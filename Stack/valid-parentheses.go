package main

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := []rune{}

	pattern := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		}
		if c == ')' || c == '}' || c == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != pattern[c] {
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
