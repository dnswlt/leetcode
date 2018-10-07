package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
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

func request(n int) {
	tr := &http.Transport{
		IdleConnTimeout: 30 * time.Second,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("http://192.168.99.144:8080/reset")
	if err != nil {
		fmt.Print("Cannot reset: ", err)
		return
	}
	ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	for i := 0; i < n; i++ {
		resp, err = client.Get("http://192.168.99.144:8080/ping")
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
	} else if len(os.Args) == 3 && os.Args[1] == "client" {
		n, _ := strconv.Atoi(os.Args[2])
		request(n)
	} else {
		fmt.Printf("Usage: %s <server|client N>\n", os.Args[0])
	}
}
