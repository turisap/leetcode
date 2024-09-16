package main

import (
	"fmt"
	"sync"
	"time"
)

type Done chan struct{}

func pipelineStage[IN comparable, OUT comparable](
	in <-chan IN,
	out chan<- OUT,
	done Done,
	cb func(a IN) OUT,
	nWorkers int,
) {

	wg := sync.WaitGroup{}
	for _ = range nWorkers {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				time.Sleep(time.Millisecond * 100)
				select {
				case v, ok := <-in:
					if !ok {
						return
					}
					out <- cb(v)
				case <-done:
					return
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}

func stageDouble(v float32) float32 {
	return v * 2
}

func stageReduce(v float32) float32 {
	return v / 5
}

func stageMultiThree(v float32) string {
	return fmt.Sprintf("%.2f", v*3)
}

const NWorkers = 5
const NJobs = 100

func main() {
	input := make([]float32, NJobs)

	for i, _ := range input {
		input[i] = float32(i)
	}

	done := make(chan struct{})
	inputChan := make(chan float32, NWorkers)
	doubleChan := make(chan float32, NWorkers)
	reduceChan := make(chan float32, NWorkers)
	resChan := make(chan string, NWorkers)

	go pipelineStage(inputChan, doubleChan, done, stageDouble, NWorkers)
	go pipelineStage(doubleChan, reduceChan, done, stageReduce, NWorkers)
	go pipelineStage(reduceChan, resChan, done, stageMultiThree, NWorkers)

	go func() {
		for _, v := range input {
			inputChan <- v
		}
		close(inputChan)
	}()

	//go func() {
	//	time.Sleep(time.Second)
	//	close(done)
	//}()

Main:
	for {
		select {
		case v, ok := <-resChan:
			fmt.Println(v)
			if !ok {
				break Main
			}
		case <-done:
			fmt.Println("done")
			break Main
		}
	}
}
