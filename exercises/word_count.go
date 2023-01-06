package exercises

import "strings"

// WordCount count the words from a phrase as asked in
// https://go.dev/tour/moretypes/23.
func WordCount(s string) map[string]int {
	var fields = strings.Fields(s)
	var countMap = make(map[string]int)
	for _, text := range fields {
		var elem, ok = countMap[text]
		if !ok {
			countMap[text] = 1
		} else {
			countMap[text] = elem + 1
		}
	}
	return countMap
}
