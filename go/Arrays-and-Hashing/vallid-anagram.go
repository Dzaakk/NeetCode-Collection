package main

func validAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	mapString := make(map[rune]int)

	for i, v := range str1 {
		mapString[v]++
		mapString[rune(str2[i])]--
	}

	for _, v := range mapString {
		if v > 0 {
			return false
		}
	}

	return true
}

// func main() {
// 	s1 := "racecar"
// 	s2 := "carrace"

// 	s3 := "jar"
// 	s4 := "jam"

// 	fmt.Println(validAnagram(s1, s2))
// 	fmt.Println(validAnagram(s3, s4))
// }
