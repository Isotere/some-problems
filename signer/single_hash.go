package main

import (
	"fmt"
	"strconv"
	"sync"
)

var mu = sync.Mutex{}

func SingleHash(in chan interface{}, out chan interface{}) {
	wg := sync.WaitGroup{}

	for data := range in {
		dataInt, _ := data.(int)
		dataStr := strconv.Itoa(dataInt)

		wg.Add(1)
		go func(data *string, out chan interface{}) {
			defer wg.Done()
			out <- singleHashCalculate(*data)
		}(&dataStr, out)
	}

	wg.Wait()
}

func singleHashCalculate(data string) string {
	wg := sync.WaitGroup{}
	wg.Add(2)

	firstVal := ""
	go func(data *string, result *string) {
		defer wg.Done()

		*result = DataSignerCrc32(*data)
	}(&data, &firstVal)

	secondVal := ""
	go func(data *string, result *string) {
		defer wg.Done()

		mu.Lock()
		md5Hash := DataSignerMd5(*data)
		mu.Unlock()

		*result = DataSignerCrc32(md5Hash)
	}(&data, &secondVal)

	wg.Wait()

	return fmt.Sprintf("%s~%s", firstVal, secondVal)
}
