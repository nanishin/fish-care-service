# fish_care_service
## To support happy life of pet fish at home

R.I.P - Gyeyomi (Blue Betta Fish Male, Late October 2016 - 16 July 2017)

![](/media/Gyeyomi.jpg)

Gyeyomi was a first pet fish brought by my first daughter.

She got Gyeyomi from biology class of elementry school.

I learned a lot of knowledge to take care of happy life of Gyeyomi.

Unfortunately he passed away to Yong-gung(i.e. Dragon palace, fish version of heaven).

Now new Guppy family started their life in the same tank.

To memorize Gyeyomi and share knowledge, this repository is created.

----

## Galup Collaboration
All things are possible by Galup collaboration.

## Arm Docker Container
Basically fish_care_service is packed to arm docker container.

The way is come from resion.io blog.

[Building ARM containers on any x86 machine, even DockerHub](https://resin.io/blog/building-arm-containers-on-any-x86-machine-even-dockerhub/)

### Arduino Control via Firmata Protocol
For physical control of fish tank, arduino gpio is used.

fish_care_service control the arduino via firmata protocol.
#### Micro Servo Motor
Support auto feeding
#### Adafruit Neopixel LED Strip
Support light
#### Power Relay
Support water pumped filter
#### Dfrobot PH Meter Sensor
Support PH value measurement
#### Dfrobot EC Meter Sensor
Support EC value measurement

### Telegram Chat Bot Control
For remote control&report, telegram bot is used.

