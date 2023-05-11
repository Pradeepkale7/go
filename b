/////blink led 
import RPi.GPIO as GPIO
import time
GPIO.setmode(GPIO.BOARD)
GPIO.setwarnings(False)
ledPin = 12
GPIO.setup(ledPin, GPIO.OUT)

for i in range(5):
 print("LED turning on.")
 GPIO.output(ledPin, GPIO.HIGH)
 time.sleep(0.5)
 print("LED turning off.")
 GPIO.output(ledPin, GPIO.LOW) 
 time.sleep(0.5)
 
 
 
 /////////////toggle two leds
 
import RPi.GPIO as GPIO
import time
GPIO.setmode(GPIO.BOARD)
GPIO.setwarnings(False)
ledPin1 = 14
ledPin2 = 15
GPIO.setup(ledPin1, GPIO.OUT)
GPIO.setup(ledPin2, GPIO.OUT)
time.sleep(1)
GPIO.output(ledPin1, True) 
GPIO.output(ledPin2, False) 
time.sleep(1)
GPIO.output(ledPin1, False) 
GPIO.output(ledPin2, True) 
time.sleep(1)
GPIO.output(ledPin1, True) 
GPIO.output(ledPin2, False) 
time.sleep(1)
GPIO.cleanup()

//////////on of buzzer

import RPi.GPIO as GPIO
from time import sleep

GPIO.setmode(GPIO.BOARD)
GPIO.setwarnings(False)

BUZZER=18
GPIO.setup(BUZZER, GPIO.OUT,initial=GPIO.LOW)
while True: 
GPIO.output(BUZZER,True)
sleep(1)
GPIO.output(BUZZER,False)
sleep(1)

