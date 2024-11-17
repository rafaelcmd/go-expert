package main

import (
	"fmt"
	"time"
)

type Message struct {
	id  int
	msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	i := 0

	// RabbitMQ
	go func() {
		for {
			i++
			time.Sleep(time.Second * 2)
			msg := Message{id: i, msg: "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()

	// Kafka
	go func() {
		for {
			i++
			time.Sleep(time.Second * 1)
			msg := Message{id: i, msg: "Hello from Kafka"}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg := <-c1: // rabbitmq

			fmt.Printf("Received from RabbitMQ: %d - %s\n", msg.id, msg.msg)
		case msg := <-c2: // kafka

			fmt.Printf("Received from Kafka: %d - %s\n", msg.id, msg.msg)
		case <-time.After(time.Second * 3):
			println("timeout")
		}
	}
}
