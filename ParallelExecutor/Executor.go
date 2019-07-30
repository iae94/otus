package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

// Параллельное исполнение
// Сделать функцию для параллельного выполнения N заданий.
// Принимает на вход слайс с заданиями `[]func()error`, число заданий которые можно выполнять параллельно `N` и максимальное число ошибок после которого нужно приостановить обработку.
// Учесть что задания могу выполняться разное время.

func Executor(funcs []func()error, N int, maxErrs int) error{

	if len(funcs) == 0 {
		return errors.New("tasks slice is empty")
	}
	if N == 0{
		return errors.New("N cannot be zero")
	}
	if maxErrs == 0{
		return errors.New("maxErrs cannot be zero")
	}
	if N > len(funcs){
		N = len(funcs)
	}
	if maxErrs > len(funcs){
		maxErrs = len(funcs)
	}

	var wg sync.WaitGroup
	var mut sync.Mutex
	resultChannel := make(chan struct{}, N)
	defer close(resultChannel)

	wg.Add(len(funcs))
	errCount := 0

	for _, fun := range funcs{
		if errCount >= maxErrs{
			err := errors.New(fmt.Sprintf("Max count of errors: %v\n", maxErrs))
			log.Println(err)
			return err
		}

		resultChannel <- struct{}{}
		go func(fun func() error) {
			err := fun()
			if err != nil {
				mut.Lock()
				errCount++
				mut.Unlock()
			}
			wg.Done()
			<- resultChannel
		}(fun)
	}
	wg.Wait()
	return nil
}
