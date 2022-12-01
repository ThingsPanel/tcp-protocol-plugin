package client

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	Client mqtt.Client
)

func NewMQTTClient(server, username string) mqtt.Client {
	opts := mqtt.NewClientOptions().
		AddBroker(server).
		//SetClientID(fmt.Sprintf("tcp-plugin-%s", username)).
		SetUsername(username).
		SetAutoReconnect(true)
	client := mqtt.NewClient(opts)

	return client
}
