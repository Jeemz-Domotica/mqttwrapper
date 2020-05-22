package mqttcontroller

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	urlString   = os.Getenv("MQTT_URL")
	topicString = os.Getenv("MQTT_TOPIC")
)

const (
	GracePeriod = 5 * time.Second
)

// GetUri returns the *url.urlString set in the env vars
func GetUri() *url.URL {
	urlObj, err := url.Parse(urlString)

	// Throw panic in case of error
	if err != nil {
		panic(err)
	}

	return urlObj
}

func GetTopic() string {
	return topicString
}

// CreateClient creates an MQTT client using supplied params.
func CreateClient(clientId string, uri *url.URL) *mqtt.Client {

	// Create client options first
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetClientID(clientId)

	// Create the client
	client := mqtt.NewClient(opts)

	// Generate connection token for MQTT broker
	token := client.Connect()

	// Repeatedly try to connect to broker. Do this with a linearly increasing time between connections.
	connections := 0
	for !token.WaitTimeout(time.Duration(connections+1) * GracePeriod) {
		connections += 1

		// If timeout duration becomes too long, handle structural connection issue.
		if time.Duration(connections+1)*GracePeriod > time.Minute {
			log.Fatal(`MQTT client could not be created. Broker connection could not be established.`)
		}

	}

	return &client
}
