package wifi

import (
	"machine"
)

func SPIConfigure(spiConfig *machine.SPIConfig) error {
	return machine.NINA_SPI.Configure(*spiConfig)
}
