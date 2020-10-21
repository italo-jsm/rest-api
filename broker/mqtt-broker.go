package broker

import (
	//import the Paho Go MQTT library
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type mqttClient struct {
	opts *MQTT.ClientOptions
	c    MQTT.Client
}

var client mqttClient

//Publish publishes a message to a topic
func Publish(command string) {
	createClient(&client)
	client.c.Publish("sub", 0, false, command)
	client.c.Disconnect(250)
}

func createClient(cli *mqttClient) {
	if cli.opts == nil {
		opts := MQTT.NewClientOptions().AddBroker("tcp://192.168.0.15:1883")
		opts.SetClientID("go-sample")
		c := MQTT.NewClient(opts)
		cli.opts = opts
		cli.c = c
	}
	if !cli.c.IsConnected(){
		connectClient(cli.c)
	}
}

func connectClient(client MQTT.Client){
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
