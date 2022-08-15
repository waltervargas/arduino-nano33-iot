package wifi

import (
	"errors"
	"machine"
	"time"

	"tinygo.org/x/drivers/wifinina"
)

const (
	StatusNoShield       = wifinina.StatusNoShield
	StatusIdle           = wifinina.StatusIdle
	NoSSIDAvail          = wifinina.StatusNoSSIDAvail
	StatusScanCompleted  = wifinina.StatusScanCompleted
	StatusConnected      = wifinina.StatusConnected
	StatusConnectFailed  = wifinina.StatusConnectFailed
	StatusConnectionLost = wifinina.StatusConnectionLost
	StatusDisconnected   = wifinina.StatusDisconnected
)

var (
	errSetPassphrase = errors.New("error setting passphrase for ssid")
	errGetIPAddress  = errors.New("error getting IP address")
)

type Wifi struct {
	ssid    string
	pass    string
	spi     *machine.SPIConfig
	adaptor *wifinina.Device
}

// SPIConfigure is intented to configure SPI interface
func (w *Wifi) spiConfigure() error {
	return machine.NINA_SPI.Configure(*w.spi)
}

// adaptoprConnect configures the adaptor and sets the SSID and Passphrase, this
// operation triggers the connection attempt to the network
func (w *Wifi) adaptorConnect() error {
	w.adaptor.Configure()

	// the wifi coprocessor needs some time to be ready
	// TODO: This code smells
	time.Sleep(2 * time.Second)

	// set the SSID and passhprase on the adaptor
	// this triggers the connection attempt to the network
	err := w.adaptor.SetPassphrase(w.ssid, w.pass)
	if err != nil {
		return errSetPassphrase
	}

	return nil
}

// GetConnectionStatus returns
func (w *Wifi) GetConnectionStatus() (wifinina.ConnectionStatus, error) {
	return w.adaptor.GetConnectionStatus()
}

// GetIpAddress
func (w *Wifi) GetIPAddress() (string, error) {
	ip, _, _, err := w.adaptor.GetIP()
	if err != nil {
		return "", errGetIPAddress
	}

	return ip.String(), nil
}

// New creates a *Wifi instance
func New(ssid, pass string) *Wifi {
	return &Wifi{
		ssid: ssid,
		pass: pass,
		spi: &machine.SPIConfig{
			// frequency on 8 Megahertz
			// chip support higher frequencies but 8 MHz is ok
			// for general use cases.
			Frequency: 8 * 1e6,
			SDO:       machine.NINA_SDO,
			SDI:       machine.NINA_SDI,
			SCK:       machine.NINA_SCK,

			Mode:     0,
			LSBFirst: false,
		},
		adaptor: &wifinina.Device{
			SPI:   machine.NINA_SPI,
			CS:    machine.NINA_CS,
			ACK:   machine.NINA_ACK,
			GPIO0: machine.NINA_GPIO0,
			RESET: machine.NINA_RESETN,
		},
	}
}

// Connect configures the SPI interface, then configures the adaptor with the
// given `ssid` and `pass`
//
// Returns a *Wifi reference that can be used to monitor the connection status via
// wifi.GetConnectionSatus()
func Connect(ssid, pass string) (*Wifi, error) {
	w := New(ssid, pass)

	// Configure SPI interface
	err := w.spiConfigure()
	if err != nil {
		return nil, err
	}

	err = w.adaptorConnect()
	if err != nil {
		return nil, err
	}

	return w, nil
}
