package wifi

import (
	"machine"
	"testing"
)

func TestSPIConfigure(t *testing.T) {
	spiConfig := machine.SPIConfig{
		// frequency on 8 Megahertz
		// chip support higher frequencies but 8 MHz is ok
		// for general use cases.
		Frequency: 8 * 1e6,
		SDO:       machine.NINA_SDO,
		SDI:       machine.NINA_SDI,
		SCK:       machine.NINA_SCK,

		Mode:     0,
		LSBFirst: false,
	}

	err := SPIConfigure(&spiConfig)

	if err != nil {
		t.Errorf("error configuring SPI: %s", err)
	}
}
