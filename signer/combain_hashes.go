package main

import (
	"sort"
	"strings"
)

func CombineResults(in chan interface{}, out chan interface{}) {
	result := make([]string, 0)

	for data := range in {
		dataStr, _ := data.(string)
		result = append(result, dataStr)
	}

	sort.Strings(result)

	out <- strings.Join(result, "_")
}
