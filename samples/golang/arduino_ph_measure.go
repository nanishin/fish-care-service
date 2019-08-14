package main

import (
	"fmt"
	"github.com/nanishin/goduino"
	"time"
)

var ph_analog_pin = 2 // A2 - 16 by ino.digitalPin in case of Analog

//var arduino = goduino.New("myArduino", "/dev/ttyUSB0")
var arduino = goduino.New("myArduino", "/dev/tty.usbmodem1411")
func main() {
	err := arduino.Connect()
	defer arduino.Disconnect()
	if err != nil {
		fmt.Println(err)
	}
	arduino.PinMode(13, goduino.Output) // Embeded LED
    //arduino.PinMode(ph_analog_pin, goduino.Analog) // No need to set mode

	for {
		arduino.DigitalWrite(13, 1)
        arduino.PhReport(ph_analog_pin)
        time.Sleep(1*time.Second) // 1 sec delay to consider measurement time
        fmt.Printf("Current pH value : "+arduino.PhValue()+"\n")
		arduino.DigitalWrite(13, 0)
		arduino.Delay(time.Millisecond * 1500)
	}
}

