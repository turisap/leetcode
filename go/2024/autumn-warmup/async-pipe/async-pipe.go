package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"sync"
	"time"
)

const workN = 3

func workerPoolPipelineStage[IN any, OUT any](input <-chan IN, output chan<- OUT, process func(IN) OUT, numWorkers int) {
	defer close(output)

	wg := sync.WaitGroup{}
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for data := range input {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(30)))
				output <- process(data)
			}
		}()
	}

	wg.Wait()
}

func asynchronousPipeline(input *csv.Reader) {
	parseInputCh := make(chan []string, workN*3)
	convertInputCh := make(chan Record, workN*3)
	encodeInputCh := make(chan Record, workN*3)
	outputCh := make(chan []byte, workN*3)

	done := make(chan struct{})

	numWorkers := 2
	go workerPoolPipelineStage(parseInputCh, convertInputCh, parse,
		numWorkers)
	go workerPoolPipelineStage(convertInputCh, encodeInputCh,
		convert, numWorkers)
	go workerPoolPipelineStage(encodeInputCh, outputCh, encode,
		numWorkers)

	go func() {
		for data := range outputCh {
			fmt.Println(string(data))
		}
		close(done)
	}()

	input.Read()
	for {
		rec, err := input.Read()
		if err == io.EOF {
			close(parseInputCh)
			break
		}
		if err != nil {
			panic(err)
		}

		parseInputCh <- rec
	}

	<-done
}
