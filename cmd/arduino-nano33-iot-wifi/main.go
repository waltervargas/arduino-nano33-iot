package main

import (
	"os"
	"time"

	"github.com/waltervargas/arduino-nano33-iot/wifi"
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
}
