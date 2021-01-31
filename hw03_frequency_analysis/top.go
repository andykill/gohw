package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	counter := map[string]int{}
	stringSlice := strings.Fields(str)
	for _, word := range stringSlice {
		_, ok := counter[word]
		if ok {
			counter[word]++
		} else {
			counter[word] = 1
		}
	}
	type strCount struct {
		Str   string
		Count int
	}

	items := []strCount{}
	for k, v := range counter {
		items = append(items, strCount{k, v})
	}

	sort.Slice(items, func(i, j int) bool {
		itemI := items[i]
		itemJ := items[j]

		if itemI.Count == itemJ.Count {
			return itemI.Str < itemJ.Str
		}

		return itemI.Count > itemJ.Count
	})

	result := []string{}
	for _, item := range items {
		if len(result) == 10 {
			break
		}
		result = append(result, item.Str)
	}

	return result
}
