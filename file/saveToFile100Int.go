package file

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func saveToFile100Int() {
	data := make(chan int)
	done := make(chan bool)
	var wg sync.WaitGroup

	logFile, _ := os.Create("log.txt")
	defer logFile.Close()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			data <- rand.Intn(200)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		done <- true
	}()

	for {
		select {
		case value := <-data:
			_, wariteErr := fmt.Fprintln(logFile, value)
			if wariteErr != nil {
				done <- false
			}
		case signal := <-done:
			if signal {
				fmt.Println("Everything is good, all done!")
				return
			} else {
				fmt.Println("Something went wrong")
				return
			}
		}
	}
}
