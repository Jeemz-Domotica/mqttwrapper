# MQTT helper package in Go
In order to streamline the coding process of MQTT message sending we abstracted a few processes away. This package assumes 
that the `MQTT_URL` and `MQTT_TOPIC` environment variables have been set in the Docker container running this code.

## Creating a client
After importing the package a `client` can be created as follows:

```go
// Get mqtt URI and topic
uri := mqttController.GetUri()
topic := mqttController.GetTopic()
name := `clientNameHere`

// Create new subscriber client
client := *mqttController.CreateClient(name, uri)
```

## Publishing a message
After getting a `client` you are ready to start publishing!

```go
client.Publish(topic, 0, false, "Ping!")
```

## Subscribing to a topic
To receive a message from an MQTT broker use:

```go
// Subscribe to MQTT and print consumed messages
client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
    // Do something with each consumed message
    fmt.Printf("Consumed: [%s] %s\n", msg.Topic(), string(msg.Payload()))
})
```

Note that the `msg.Payload()` first needs to be converted to a `string`, given that is is of type `[]byte`!

## Example usage
The `docker-compose.yaml` can be used to run a set of containers that run a simple setup using the `mqttController` package.
The code for these containers can be found in `./examples/`. This can also be used to test alterations made to this package.
