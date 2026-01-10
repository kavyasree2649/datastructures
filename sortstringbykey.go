package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	key   string
	value string
}

func main() {
	arr := []string{"92 022", "82 12", "77 13"}
	key := 2
	n := 3
	reverse := false
	compare := "numeric"

	pairs := make([]Pair, n)

	for i := 0; i < n; i++ {
		pairs[i] = Pair{
			key:   extractKeyFromString(arr[i], key),
			value: arr[i],
		}
	}

	if compare == "numeric" {
		sort.Slice(pairs, func(i, j int) bool {
			a, _ := strconv.Atoi(pairs[i].key)
			b, _ := strconv.Atoi(pairs[j].key)
			if reverse {
				return a > b
			}
			return a < b
		})
	}

	if compare == "lexico" {
		sort.Slice(pairs, func(i, j int) bool {

			if reverse {
				return pairs[i].key > pairs[j].key
			}
			return pairs[i].key < pairs[j].key
		})
	}

	result := make([]string, 0)

	for i := 0; i < n; i++ {
		result = append(result, pairs[i].value)
	}

	fmt.Println(result)
}

func extractKeyFromString(inputString string, key int) string {

	slice := strings.Fields(inputString)

	if len(slice) >= key {
		return slice[key-1]
	}
	return ""
}
