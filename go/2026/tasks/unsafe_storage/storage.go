package main

import "sync"

type Connection interface {
	Connect()

	Disconnect()

	Send(req string) (string, error)
}

type ConnectionCreator interface {
	NewConnection() (Connection, error)
}

type Saver interface {
	Save(data string)
}

func SendAndSave(creator ConnectionCreator, saver Saver, requests []string, maxConn int) {
	// 	jobs worker
	wg := &sync.WaitGroup{}
	jobs := make(chan string, maxConn)
	mutex := &sync.Mutex{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, req := range requests {
			jobs <- req
		}

		close(jobs)
	}()

	for i := 0; i < maxConn; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			conn, err := creator.NewConnection()
			if err != nil {
				return
			}

			conn.Connect()
			defer conn.Disconnect()

			for job := range jobs {
				val, err := conn.Send(job)
				if err != nil {
					continue
				}

				mutex.Lock()
				saver.Save(val)
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
}
