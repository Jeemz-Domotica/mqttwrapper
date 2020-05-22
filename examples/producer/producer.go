package main

import (
	"fmt"
	"github.com/Jeemz-Domotica/mqttwrapper"
	"os"
	"time"
)

// Sample producer that produces a PING every 10 seconds.
func main() {
	// Wait a couple of seconds to ensure MQTT is initialized.
	time.Sleep(mqttwrapper.GracePeriod)

	// Get URL and topic of MQTT.
	uri := mqttwrapper.GetUri()
	topic := mqttwrapper.GetTopic()
	producerName := os.Getenv(`PUB_NAME`)

	// Create MQTT client.
	client := *mqttwrapper.CreateClient(producerName, uri)

	// Publish message every 10 seconds.
	timer := time.NewTicker(10 * time.Second)
	for range timer.C {
		client.Publish(topic, 0, false, "Ping!")
		fmt.Println("Produced message!")
	}
}
