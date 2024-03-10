package main

import (
	"strconv"
	"strings"
	"sync"
)

func MultiHash(in chan interface{}, out chan interface{}) {
	wg := sync.WaitGroup{}

	for data := range in {
		dataStr, _ := data.(string)

		wg.Add(1)
		go func(data *string, out chan interface{}) {
			defer wg.Done()
			out <- multiHashCalculate(*data)
		}(&dataStr, out)
	}

	wg.Wait()
}

func multiHashCalculate(data string) string {
	wg := sync.WaitGroup{}
	wg.Add(6)

	result := make([]string, 6)
	for i := 0; i < 6; i++ {
		go func(i int, data *string, result []string) {
			defer wg.Done()

			result[i] = DataSignerCrc32(strconv.Itoa(i) + *data)
		}(i, &data, result)
	}

	wg.Wait()

	return strings.Join(result, "")
}
