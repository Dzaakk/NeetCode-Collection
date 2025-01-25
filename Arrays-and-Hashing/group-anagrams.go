package main

import (
	"fmt"
	"strconv"
	"strings"
)

func groupAnagrams(words []string) [][]string {
	anagramGroups := make(map[string][]string)

	for _, word := range words {
		key := generateRuneKey(word)

		anagramGroups[key] = append(anagramGroups[key], word)
	}

	var result [][]string
	for _, group := range anagramGroups {
		result = append(result, group)
	}

	return result
}

func generateRuneKey(word string) string {
	runeCount := make([]int, 26)

	for _, r := range word {
		runeCount[r-'a']++
	}

	var keyBuilder strings.Builder
	for _, count := range runeCount {
		keyBuilder.WriteString(strconv.Itoa(count))
		keyBuilder.WriteRune('#')
	}

	return keyBuilder.String()
}

func main() {
	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}

	fmt.Println(groupAnagrams(strs))
}
