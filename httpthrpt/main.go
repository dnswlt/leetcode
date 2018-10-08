package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func serve() {
	countChan := make(chan int, 100)
	resetChan := make(chan bool)
	timerChan := make(chan bool)
	go func() {
		for {
			timerChan <- true
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		total := 0
		started := time.Now()
		lastMessageReceived := started
		for {
			select {
			case b := <-countChan:
				total += b
				lastMessageReceived = time.Now()
			case <-resetChan:
				total = 0
				started = time.Now()
			case <-timerChan:
				elapsed := lastMessageReceived.Sub(started)
				fmt.Printf("Received %d messages in %dms (%.1f msg/s)\n", total, elapsed.Nanoseconds()/1000000, float64(total)/elapsed.Seconds())
			}
		}
	}()
	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ACK")
		resetChan <- true
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
		countChan <- 1
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func reset(host string, port int) {
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/reset", host, port))
	if err != nil {
		fmt.Print("Cannot reset: ", err)
		return
	}
	ioutil.ReadAll(resp.Body)
	resp.Body.Close()
}

func request(n int, host string, port int) {
	tr := &http.Transport{
		IdleConnTimeout: 30 * time.Second,
	}
	client := &http.Client{Transport: tr}
	for i := 0; i < n; i++ {
		resp, err := client.Get(fmt.Sprintf("http://%s:%d/ping", host, port))
		if err != nil {
			fmt.Print("Cannot ping: ", err)
			return
		}
		ioutil.ReadAll(resp.Body)
		resp.Body.Close()
	}
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "server" {
		serve()
	} else if len(os.Args) == 6 && os.Args[1] == "client" {
		host := os.Args[2]
		port, _ := strconv.Atoi(os.Args[3])
		numClients, _ := strconv.Atoi(os.Args[4])
		numRequests, _ := strconv.Atoi(os.Args[5])
		reset(host, port)
		var wg sync.WaitGroup
		wg.Add(numClients)
		for i := 0; i < numClients; i++ {
			go func() {
				defer wg.Done()
				request(numRequests, host, port)
			}()
		}
		wg.Wait()
	} else {
		fmt.Printf("Usage: %s <server|client host port num-clients num-requests>\n", os.Args[0])
	}
}
