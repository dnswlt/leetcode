package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

const (
	exchangeName      = "load"
	rabbitURL         = "amqp://192.168.99.144:5672/"
	messageRoutingKey = "message"
	resetRoutingKey   = "reset"
	numMessages       = 512
)

func newMessage(body string) amqp.Publishing {
	return amqp.Publishing{
		DeliveryMode: amqp.Transient,
		Timestamp:    time.Now(),
		ContentType:  "text/plain",
		Body:         []byte(body),
	}
}

func send(ch *amqp.Channel, n int) {
	err := ch.Publish(exchangeName, resetRoutingKey, false, false, newMessage("reset"))
	if err != nil {
		log.Fatal("Could not reset: ", err)
	}
	started := time.Now()
	var data strings.Builder
	for i := 0; i < numMessages; i++ {
		data.WriteRune('a')
	}
	msg := newMessage(data.String())
	for i := 0; i < n; i++ {
		err := ch.Publish(exchangeName, messageRoutingKey, false, false, msg)
		if err != nil {
			log.Fatal("Could not publish message: ", err)
		}
	}
	elapsed := time.Since(started)
	log.Printf("Sent %d messages in %dms\n", n, elapsed.Nanoseconds()/1000000)
}

func bind(ch *amqp.Channel, name string, routingKey string) {
	err := ch.QueueBind(name, routingKey, exchangeName, false, nil)
	if err != nil {
		log.Fatal("Could not bind queue: ", err)
	}
}

func recv(ch *amqp.Channel) {
	q, err := ch.QueueDeclare("", false, true, true, false, nil)
	if err != nil {
		log.Fatal("Could not declare queue: ", err)
	}
	bind(ch, q.Name, resetRoutingKey)
	bind(ch, q.Name, messageRoutingKey)
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Cannot consume: ", err)
	}
	var count int
	tickChan := make(chan struct{})
	go func() {
		for {
			time.Sleep(1 * time.Second)
			tickChan <- struct{}{}
		}
	}()
	started := time.Now()
	lastMessageReceived := started
	for {
		select {
		case m, ok := <-msgs:
			if !ok {
				log.Fatal("Connection lost: ", err)
			}
			if m.RoutingKey == resetRoutingKey {
				count = 0
			} else {
				if count == 0 {
					started = time.Now()
				}
				count++
				lastMessageReceived = time.Now()
			}
		case <-tickChan:
			elapsed := lastMessageReceived.Sub(started)
			log.Printf("Received %d messages in %dms (%.1f msg/s)\n", count, elapsed.Nanoseconds()/1000000, float64(count)/elapsed.Seconds())
		}

	}

}

func main() {
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatal("Could not connect: ", err)
	} else {
		log.Print("Connected.")
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Could not get a channel", err)
	}
	defer ch.Close()
	err = ch.ExchangeDeclare(exchangeName, "topic", false, false, false, false, nil)
	if err != nil {
		log.Fatal("Could not declare topic exchange", err)
	}
	if os.Args[1] == "send" {
		n, _ := strconv.Atoi(os.Args[2])
		send(ch, n)
	} else {
		recv(ch)
	}

}
