package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/waltervargas/arduino-nano33-iot/mqttclient"
	"github.com/waltervargas/arduino-nano33-iot/wifi"
)

const (
	mqttserver = "tcp://192.168.0.118:1883"
	clientid   = "arduino-nano33-123"
)

func main() {
	// ssid and pass defined as constants at secrets.go
	w, err := wifi.Connect(ssid, pass)
	if err != nil {
		println("error connecting to wifi: " + err.Error())
		os.Exit(1)
	}

	//check the connection status and wait for the connection to be success
	for st, _ := w.GetConnectionStatus(); st != wifi.StatusConnected; {
		println("Connection status: " + st.String())
		time.Sleep(2 * time.Second)
		st, _ = w.GetConnectionStatus()
	}

	println("connected to: " + ssid)
	ipAddress, err := w.GetIPAddress()
	if err != nil {
		println(err.Error())
	}

	println("IP address: " + ipAddress)

	mqttc, err := mqttclient.Connect(mqttserver, clientid)
	if err != nil {
		println("error connecting to mqtt: " + err.Error())
	}

	println("connected to MQTT server: " + mqttserver)

	// testing sending messages to MQTT
	println("publishing messages to server: " + mqttserver)

	for {
		rand.Seed(time.Now().Unix())
		temp := rand.Intn(35)
		err = mqttc.PublishMessage("iot/temperature", fmt.Sprintf(`{value: %v}`, temp))
		if err != nil {
			println("error sending message to mqtt")
		}

		time.Sleep(5 * time.Second)
	}

}
