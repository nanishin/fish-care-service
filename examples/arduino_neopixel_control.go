package main

import (
	"fmt"
	"github.com/nanishin/goduino"
	"time"
)

var light_neopixel_pin = 2
var light_neopixel_num = 36

//var arduino = goduino.New("myArduino", "/dev/ttyUSB0")
var arduino = goduino.New("myArduino", "/dev/tty.usbmodem1411")
func main() {
	err := arduino.Connect()
	defer arduino.Disconnect()
	if err != nil {
		fmt.Println(err)
	}
	arduino.PinMode(13, goduino.Output) // Embeded LED
    arduino.PinMode(light_neopixel_pin, goduino.Pwm)

	for {
		arduino.DigitalWrite(13, 1)
		ArduinoLightControl("ON")
		arduino.Delay(time.Millisecond * 1500)
		ArduinoLightControl("RED")
		arduino.Delay(time.Millisecond * 1500)
		ArduinoLightControl("GREEN")
		arduino.Delay(time.Millisecond * 1500)
		ArduinoLightControl("BLUE")
		arduino.Delay(time.Millisecond * 1500)
		ArduinoLightControl("OFF")
		arduino.DigitalWrite(13, 0)
		arduino.Delay(time.Millisecond * 1500)
	}
}

var light_led_status = "OFF"
func ArduinoLightControl(state string) {
	if state == "ON" {
		// Fish House Light - 0:White, 1:Red, 2:Green, 3:Blue
		arduino.NeopixelControl(light_neopixel_pin, light_neopixel_num, 0, 1)
		fmt.Printf("Arduino Neopixel Turn On!!!\n")
	} else if state == "RED" {
		// Fish House Light - 0:White, 1:Red, 2:Green, 3:Blue
		arduino.NeopixelControl(light_neopixel_pin, light_neopixel_num, 1, 1)
		fmt.Printf("Arduino Neopixel Turn Red!!!\n")
	} else if state == "GREEN" {
		// Fish House Light - 0:White, 1:Red, 2:Green, 3:Blue
		arduino.NeopixelControl(light_neopixel_pin, light_neopixel_num, 2, 1)
		fmt.Printf("Arduino Neopixel Turn Green!!!\n")
	} else if state == "BLUE" {
		// Fish House Light - 0:White, 1:Red, 2:Green, 3:Blue
		arduino.NeopixelControl(light_neopixel_pin, light_neopixel_num, 3, 1)
		fmt.Printf("Arduino Neopixel Turn Blue!!!\n")
	} else if state == "OFF" {
		// Fish House Light - 0:White, 1:Red, 2:Green, 3:Blue
		arduino.NeopixelControl(light_neopixel_pin, light_neopixel_num, 0, 0)
		fmt.Printf("Arduino Neopixel Turn Off!!!\n")
	}
	light_led_status = state
}
