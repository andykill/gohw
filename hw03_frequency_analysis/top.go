package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"sort"
	"strings"
	"unicode"
)

func Top10(str string) []string {
	str = strings.ToLower(str)
	counter := map[string]int{}

	for _, word := range strings.FieldsFunc(str, isAccept) {
		if word == "-" {
			continue
		}
		if keyInMap(word, counter) {
			counter[word]++
		}
	}
	keys := sortMao(counter)
	result := []string{}
	for _, k := range keys {
		if len(result) == 10 {
			break
		}
		result = append(result, k)
	}

	return result
}

func keyInMap(str string, strMap map[string]int) bool {
	for v := range strMap {
		if v == str {
			return true
		}
	}
	return false
}

func isAccept(r rune) bool {
	return r != 45 && !unicode.IsLetter(r) && !unicode.IsNumber(r)
}

func sortMao(counter map[string]int) []string {
	keys := make([]string, 0, len(counter))
	for k := range counter {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
