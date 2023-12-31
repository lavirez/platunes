package main

import (
    "fmt"
    "log"
    "os"

    "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
    if err != nil {
        panic(fmt.Sprintf("%s: %s", msg, err))
    }
}

func main() {
    amqpUrl := os.Getenv("AMQP_URL")
    if amqpUrl == nil { 
        amqpUrl ="amqp://rabbitmq:brabbit@rabbitmq-server:5672/"
    }

    conn, err = amqp.Dial(amqpUrl)
    failOnError(err, "Error connecting to the broker")
    
    defer conn.Close()

    ch, err = conn.Channel()
    failOnError(err, "Failed to Open the Channel")

    defer ch.Close()

    exhcangeName := "user_updates"
    bindingKey := "user.profile"

    err = ch.ExchangeDeclare(
        exchangeName,   
        "topic",       
        true,           
        false,
        false,
        false,
        nil,
    )
    failOnError(err, "Error creating the exchange")

    queue, err := ch.QueueDeclare(
        "", //random but unique by pkg
        true,
        false,
        false,
        false,
        nil,
    )
    failOnError(err, "Error creating the queue")

    msgs, err := ch.Consume(
        q.Name,
        "",
        false,
        false,
        false,
        false,
        nil,
    )
    failOnError(err, "Failed to register as a consumer")

    forever := make(chan bool)

    go func() { 
        for d := range msgs {
            log.Printf("Received messages: %s", d.body)

            d.Ack(false)
    }

    fmt.Prinln("Sevices listening for events...")

    <-forever
}





