package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type message struct {
	id   int64
	data []byte
}

func send(ch chan<- message, numMessages int) {
	for i := 0; i < numMessages; i++ {
		data := make([]byte, 512)
		ch <- message{int64(i), data}
	}
	close(ch)
}

func recv(ch <-chan message, wg *sync.WaitGroup) {
	defer wg.Done()
	cnt := 0
	started := time.Now()
	for range ch {
		cnt++
	}
	elapsed := time.Since(started)
	fmt.Printf("Received %d messages in %.3fms (%.1f msg/s)\n", cnt, elapsed.Seconds()*1000, float64(cnt)/elapsed.Seconds())
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <NumMessages>\n", os.Args[0])
		return
	}
	numMessages, _ := strconv.Atoi(os.Args[1])
	ch := make(chan message, 1024)
	var wg sync.WaitGroup
	wg.Add(1)
	go send(ch, numMessages)
	go recv(ch, &wg)
	wg.Wait()
}
