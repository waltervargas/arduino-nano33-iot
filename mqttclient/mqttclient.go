package mqttclient

import (
	"tinygo.org/x/drivers/net/mqtt"
)

type Client struct {
	server     string
	id         string
	mqttClient mqtt.Client
}

func New(server, id string) Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(server).SetClientID(id)

	return Client{
		server:     server,
		id:         id,
		mqttClient: mqtt.NewClient(opts),
	}
}

func (c Client) connect() error {
	if token := c.mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func Connect(server, id string) (Client, error) {
	c := New(server, id)
	err := c.connect()
	if err != nil {
		return Client{}, err
	}

	return c, nil
}

func (c Client) PublishMessage(path, message string) error {
	token := c.mqttClient.Publish(path, 1, false, message)
	token.Wait()
	if token.Error() != nil {
		return token.Error()
	}

	return nil
}
