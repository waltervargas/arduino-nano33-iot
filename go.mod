module github.com/waltervargas/arduino-nano33-iot

go 1.18

replace device/arm => /usr/local/lib/tinygo/src/device/arm

replace device/sam => /usr/local/lib/tinygo/src/device/sam

replace internal/reflectlite => /usr/local/lib/tinygo/src/internal/reflectlite

replace internal/task => /usr/local/lib/tinygo/src/internal/task

replace machine => /usr/local/lib/tinygo/src/machine

replace os => /usr/local/lib/tinygo/src/os

replace reflect => /usr/local/lib/tinygo/src/reflect

replace runtime => /usr/local/lib/tinygo/src/runtime

replace runtime/interrupt => /usr/local/lib/tinygo/src/runtime/interrupt

replace runtime/volatile => /usr/local/lib/tinygo/src/runtime/volatile

replace sync => /usr/local/lib/tinygo/src/sync

replace testing => /usr/local/lib/tinygo/src/testing

require machine v0.0.0-00010101000000-000000000000

require github.com/eclipse/paho.mqtt.golang v1.2.0 // indirect

require (
	device/arm v0.0.0-00010101000000-000000000000 // indirect
	device/sam v0.0.0-00010101000000-000000000000 // indirect
	runtime/interrupt v0.0.0-00010101000000-000000000000 // indirect
	runtime/volatile v0.0.0-00010101000000-000000000000 // indirect
	tinygo.org/x/drivers v0.21.0
)
