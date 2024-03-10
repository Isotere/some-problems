package main

import (
	"sync"
)

func ExecutePipeline(jobs ...job) {
	if len(jobs) == 0 {
		return
	}

	// Инициализируем входные-выходные каналы на 1 больше, чем кол-во задач
	channels := make([]chan interface{}, len(jobs)+1)
	for i := range channels {
		channels[i] = make(chan interface{}, 1)
	}

	wg := sync.WaitGroup{}

	inChI, outChI := 0, 1
	for _, currJob := range jobs {
		wg.Add(1)
		go func(in chan interface{}, out chan interface{}) {
			defer wg.Done()
			// Закрываем выходные каналы после выполнения функций, чтобы все возможные range, select закончились
			defer close(out)

			currJob(in, out)
		}(channels[inChI], channels[outChI])

		inChI++
		outChI++
	}

	wg.Wait()
}
