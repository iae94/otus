package main

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestExecutor(t *testing.T) {

	successWorkerFast := func() error{
		fmt.Printf("Start FastWorker\n")
		time.Sleep(1 * time.Second)
		fmt.Printf("End FastWorker\n")
		return nil
	}
	successWorkerLong := func() error{
		fmt.Printf("Start LongWorker\n")
		time.Sleep(5 * time.Second)
		fmt.Printf("End LongWorker\n")
		return nil
	}


	errWorkerFast := func() error{
		fmt.Printf("Start FastErrWorker\n")
		time.Sleep(2 * time.Second)
		fmt.Printf("End FastErrWorker\n")
		return errors.New("some error")
	}

	errWorkerLong := func() error{
		fmt.Printf("Start LongErrWorker\n")
		time.Sleep(6 * time.Second)
		fmt.Printf("End LongErrWorker\n")
		return errors.New("some error")
	}

	tests := []struct {
		name	string
		funcs   []func() error
		N 		int
		maxErrs int
		wantErr bool
	}{
		{name: "Success", funcs: []func()error{successWorkerFast, errWorkerLong, errWorkerFast, successWorkerLong}, N: 2, maxErrs:4, wantErr:false},
		{name: "N > len(funcs)", funcs: []func()error{successWorkerFast, errWorkerLong, errWorkerFast, successWorkerLong}, N: 20, maxErrs: 3, wantErr:false},
		{name: "maxErrs > len(funcs)", funcs: []func()error{successWorkerFast, errWorkerLong, errWorkerFast, successWorkerLong}, N: 2, maxErrs:20, wantErr:false},

		{name: "Errors limit", funcs: []func()error{errWorkerLong, successWorkerFast, errWorkerFast, successWorkerLong, errWorkerFast, successWorkerLong, successWorkerFast, successWorkerLong}, N: 2, maxErrs:3, wantErr:true},
		{name: "Empty funcs", funcs: []func()error{}, N: 2, maxErrs:2, wantErr:true},
		{name: "N == 0", funcs: []func()error{successWorkerFast, errWorkerLong, errWorkerFast, successWorkerLong}, N: 0, maxErrs:2, wantErr:true},
		{name: "maxErrs == 0", funcs: []func()error{successWorkerFast, errWorkerLong, errWorkerFast, successWorkerLong}, N: 2, maxErrs:0, wantErr:true},

	}
	for _, tt := range tests {

		// Exec test
		t.Run(tt.name, func(t *testing.T) {
			if err := Executor(tt.funcs, tt.N, tt.maxErrs); (err != nil) != tt.wantErr {
				t.Errorf("Executor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
