package main

import (
	"fmt"
	"github.com/nanishin/goduino"
	"time"
)

var power_relay_pin = 11

//var arduino = goduino.New("myArduino", "/dev/ttyUSB0")
var arduino = goduino.New("myArduino", "/dev/tty.usbmodem1411")
func main() {
	err := arduino.Connect()
	defer arduino.Disconnect()
	if err != nil {
		fmt.Println(err)
	}
	arduino.PinMode(13, goduino.Output) // Embeded LED
    arduino.PinMode(power_relay_pin, goduino.Output)

	for {
		arduino.DigitalWrite(13, 1)
		ArduinoPumpControl("OFF")
		arduino.Delay(time.Millisecond * 1500)
		arduino.DigitalWrite(13, 0)
		ArduinoPumpControl("ON")
		arduino.Delay(time.Millisecond * 1500)
	}
}

// Default Relay NO(Nomaly Open)
var water_pump_status = "ON"
func ArduinoPumpControl(state string) {
	if state == "ON" {
		arduino.DigitalWrite(power_relay_pin, 0)
		fmt.Printf("Arduino Pump On!!!\n")
	} else if state == "OFF" {
		arduino.DigitalWrite(power_relay_pin, 1)
		fmt.Printf("Arduino Pump Off!!!\n")
	}
	water_pump_status = state
}
