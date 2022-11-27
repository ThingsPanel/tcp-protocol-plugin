package global

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"strings"
)

type MqttClient struct {
	Client mqtt.Client
}

// raw data to send
type rawData struct {
	Token  string `json:"token"`
	Values []byte `json:"values"`
}

var (
	DefaultMqttClient *MqttClient = &MqttClient{}
)

func (c *MqttClient) Init() error {
	opts := mqtt.NewClientOptions().
		AddBroker(Config.Mqtt.Addr).
		SetClientID("custom-tcp-plugin").
		SetUsername(Config.Mqtt.Username).
		SetPassword(Config.Mqtt.Password).
		SetAutoReconnect(true)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		return token.Error()
	}

	c.Client = client
	return nil
}

func (c *MqttClient) SendRawData(accessToken string, data []byte) error {
	sendData := &rawData{
		Token:  accessToken,
		Values: data,
	}
	d, err := json.Marshal(sendData)
	if err != nil {
		return err
	}
	log.Info("send data:", string(d))
	log.Info("topic", Config.Topic.PublishRawData)
	if token := c.Client.Publish(Config.Topic.PublishRawData, byte(Config.Mqtt.Qos), false, d); token.Wait() && token.Error() != nil {
		log.Info(token.Error())
		return err
	}

	return nil
}

func (c *MqttClient) Subscribe() {
	c.Client.Subscribe(Config.Topic.SubscribeRawData, byte(Config.Mqtt.Qos), func(client mqtt.Client, msg mqtt.Message) {
		log.Info("subscribe msg:", string(msg.Payload()))

		topicArr := strings.Split(msg.Topic(), "/")
		if len(topicArr) != 3 {
			log.Info("topic error")
			return
		}
		deviceToken := topicArr[2]

		log.Info("deviceToken:", deviceToken)
		receiveData := make(map[string]string)
		err := json.Unmarshal(msg.Payload(), &receiveData)
		if err != nil {
			log.Error(err)
			return
		}
		device := Devices[deviceToken]
		log.Info(device)
		if device != nil && device.Online == true {
			device.ClientConn.Write([]byte(receiveData["values"]))
		}
	})
}
