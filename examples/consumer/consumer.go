package main

import (
	"fmt"
	"github.com/Jeemz-Domotica/mqttController"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

// Consumer that prints consumed messages indefinitely
func main() {
	// Wait for the duration of the grace period
	time.Sleep(mqttController.GracePeriod)

	// Get mqtt URI and topic
	uri := mqttController.GetUri()
	topic := mqttController.GetTopic()

	// Create new subscriber client
	client := *mqttController.CreateClient("sub", uri)

	// Subscribe to MQTT and print consumed messages
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		// Do something with each consumed message
		fmt.Printf("Consumed: [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})

	// Timer that allows listening to MQTT infinitely. Prints time every 5 minutes.
	timer := time.NewTicker(5 * time.Minute)
	for t := range timer.C {
		fmt.Println("Time: ", t)
	}
}
