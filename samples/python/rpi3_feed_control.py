import RPi.GPIO as GPIO
import time
 
pin = 4 # PWM pin num 4
 
GPIO.setmode(GPIO.BCM)
GPIO.setup(pin, GPIO.OUT)
p = GPIO.PWM(pin, 50)
p.start(0)

def SetAngle(dc):
    p.ChangeDutyCycle(dc)
    time.sleep(0.3)

for dc in range(0, 11, 1):
    SetAngle(dc)

for dc in range(10, -1, -1):
    SetAngle(dc)

p.stop()
GPIO.output(pin, False)
GPIO.cleanup()
