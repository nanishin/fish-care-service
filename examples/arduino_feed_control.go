package main

import (
	"fmt"
	"github.com/nanishin/goduino"
	"time"
)

var feed_servo_pin = 3

//var arduino = goduino.New("myArduino", "/dev/ttyUSB0")
var arduino = goduino.New("myArduino", "/dev/tty.usbmodem1421")
func main() {
	err := arduino.Connect()
	defer arduino.Disconnect()
	if err != nil {
		fmt.Println(err)
	}
	arduino.PinMode(13, goduino.Output) // Embeded LED

	for {
		arduino.DigitalWrite(13, 1)
        ArduinoFeedControl()
		arduino.DigitalWrite(13, 0)
		arduino.Delay(time.Millisecond * 1500)
	}
}

func ArduinoFeedControl() {
	arduino.PinMode(feed_servo_pin, goduino.Servo)
    arduino.Delay(time.Millisecond * 10)

/*
	var pos byte
	for pos = 0; pos <= 130; pos += 1 {
		arduino.ServoWrite(feed_servo_pin, pos)
		arduino.Delay(time.Millisecond * 10)
	}

	for pos = 130; pos > 0; pos -= 1 {
		arduino.ServoWrite(feed_servo_pin, pos)
		arduino.Delay(time.Millisecond * 10)
	}
	arduino.ServoWrite(feed_servo_pin, pos)
*/
    arduino.ServoWrite(feed_servo_pin, 0)
    arduino.Delay(time.Millisecond * 1000)
    arduino.ServoWrite(feed_servo_pin, 130)
    arduino.Delay(time.Millisecond * 1000)
    arduino.ServoWrite(feed_servo_pin, 0)

    arduino.Delay(time.Millisecond * 10)

	arduino.PinMode(feed_servo_pin, goduino.Input)
}
