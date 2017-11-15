# fish-care-service
## To support happy life of pet fish at home

R.I.P - Gwiyomi (Blue Betta Fish Male, Late October 2016 - 16 July 2017)

![](/media/Gwiyomi.jpg)

Gwiyomi was a first pet fish brought by my first daughter.

She got Gwiyomi from biology class of elementry school.

I learned a lot of knowledge to take care of happy life of Gwiyomi.

Unfortunately he passed away to Yong-gung(i.e. Dragon palace, fish version of heaven).

Now new Guppy family started their life in the same tank.

To memorize Gwiyomi and share knowledge, this repository is created.

----

## Galup Collaboration
All things are possible by Galup collaboration.

## Arm Docker Container
Basically fish-care-service is packed to arm docker container.

Referenced from resin.io blog ([Building ARM containers on any x86 machine, even DockerHub](https://resin.io/blog/building-arm-containers-on-any-x86-machine-even-dockerhub/))

Referenced from [resin-io/qemu](https://github.com/resin-io/qemu)
### Arduino Control via Firmata Protocol
For physical control of fish tank, arduino gpio is used.

fish-care-service control the arduino via firmata protocol.

Referenced from arduino [firmata](https://github.com/firmata/arduino)

I made [GalupFirmata](https://github.com/nanishin/GalupFirmata) to support more sensors based on [StandardFirmata](https://github.com/firmata/arduino/tree/master/examples/StandardFirmata).
#### Micro Servo Motor
Support auto feeding

Referenced from [golang goduino](https://github.com/nanishin/goduino)
#### Power Relay
Support water pumped filter

Referenced from [golang goduino](https://github.com/nanishin/goduino)
#### Adafruit Neopixel LED Strip
Support light

Referenced from [Adafruit Neopixel](https://github.com/adafruit/Adafruit_NeoPixel)
#### Analog PH Meter Sensor
Support PH value measurement

Referenced from [dfrobot SEN0161 manual](https://www.dfrobot.com/wiki/index.php/PH_meter(SKU:_SEN0161))
#### Analog EC Meter Sensor
Support EC value measurement

Referenced from [dfrobot DFR0300 manual](https://www.dfrobot.com/wiki/index.php/Analog_EC_Meter_SKU:DFR0300)
#### Analog Liquid Leak Sensor
Support Liquid leak status

Referenced from [cooking-hacks leak sensor manual](https://www.cooking-hacks.com/water-leakage-liquid-detection-sensor-line)
### Telegram Chat Bot Control
For remote control&report, telegram chat bot is used.

Referenced from [golang telebot](https://github.com/nanishin/telebot)

