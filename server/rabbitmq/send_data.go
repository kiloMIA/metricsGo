package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const rbmq = "amqp://guest:guest@localhost:5672/"

func main() {
	conn, err := amqp.Dial(rbmq)
	if err != nil {
		log.Panicf("Could not connect to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Panicf("Could not open channel")
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"data_direct", //exchnage name
		"direct",      //type
		true,          //durable
		false,         //auto-deleted
		false,         //internal
		false,         //no-wait
		nil,           //arguments
	)
	if err != nil {
		log.Panicf("Could not create exchange, %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"data_direct", //exchange
		"metrics",     //routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("a"),
		})

	if err != nil {
		log.Panicf("Could not publish message, %v", err)
	}
}
