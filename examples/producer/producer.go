package main

import (
	"fmt"
	"github.com/Jeemz-Domotica/mqttController"
	"os"
	"time"
)

// Sample producer that produces a PING every 10 seconds.
func main() {
	// Wait a couple of seconds to ensure MQTT is initialized.
	time.Sleep(mqttController.GracePeriod)

	// Get URL and topic of MQTT.
	uri := mqttController.GetUri()
	topic := mqttController.GetTopic()
	producerName := os.Getenv(`PUB_NAME`)

	// Create MQTT client.
	client := *mqttController.CreateClient(producerName, uri)

	// Publish message every 10 seconds.
	timer := time.NewTicker(10 * time.Second)
	for range timer.C {
		client.Publish(topic, 0, false, "Ping!")
		fmt.Println("Produced message!")
	}
}
