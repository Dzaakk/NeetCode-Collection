package main

// O(m*n)
func groupAnagrams(words []string) [][]string {
	mapGroup := make(map[[26]int][]string)

	for _, s := range words {
		var count [26]int
		for _, c := range s {
			count[c-'a']++
		}
		mapGroup[count] = append(mapGroup[count], s)

	}

	result := make([][]string, 0, len(mapGroup))
	for _, group := range mapGroup {
		result = append(result, group)
	}

	return result
}

// func groupAnagrams(words []string) [][]string {
// 	anagramGroups := make(map[string][]string)

// 	for _, word := range words {
// 		key := generateRuneKey(word)

// 		anagramGroups[key] = append(anagramGroups[key], word)
// 	}

// 	var result [][]string
// 	for _, group := range anagramGroups {
// 		result = append(result, group)
// 	}

// 	return result
// }

// func generateRuneKey(word string) string {
// 	runeCount := make([]int, 26)

// 	for _, r := range word {
// 		runeCount[r-'a']++
// 	}

// 	var keyBuilder strings.Builder
// 	for _, count := range runeCount {
// 		keyBuilder.WriteString(strconv.Itoa(count))
// 		keyBuilder.WriteRune('#')
// 	}

// 	return keyBuilder.String()
// }

// func main() {
// 	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}

// 	fmt.Println(groupAnagrams(strs))
// }
