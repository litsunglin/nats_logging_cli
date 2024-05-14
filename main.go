package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func onMessage(msg *nats.Msg) {
	fmt.Printf("> %s\n", string(msg.Data))
}

func main() {

	c := tls.Config{InsecureSkipVerify: true}
	
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL, nats.Secure(&c))
	if err != nil {
		log.Fatalf("Error connecting to NATS server: %s", err)
	} else {
		log.Printf("Connected to NATS server: %s",nc.ConnectedAddr())
	}
	defer nc.Close()

	// Subscribe to the channel
	sub, err := nc.Subscribe("debug_logging", onMessage)
	if err != nil {
		log.Fatalf("Error subscribing to channel: %s", err)
	} else {
		log.Printf("Subscribed to channel: %s", sub.Subject)
	}
	defer sub.Unsubscribe()

	// Wait for messages indefinitely
	nc.Flush()
	nc.FlushTimeout(5 * time.Second) // Adjust timeout as needed

	// Keep the program running
	select {}
}
