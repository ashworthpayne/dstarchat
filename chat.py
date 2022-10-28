#!/usr/bin/python3
import serial
from threading import Thread 

serialPort = serial.Serial(port='/dev/ttyUSB1', baudrate=9600)

class Receiver(Thread):
    def __init__(self, serialPort): 
        Thread.__init__(self)
        self.serialPort = serialPort 

    def run(self):
        text = ""
        while (text != "exit\n"):
            text = serialPort.readline()
            print ("\n machine1: " + text)

class Sender(Thread):
    def __init__(self, serialPort):
        Thread.__init__(self)
        self.serialPort = serialPort 

    def run(self):
        text = ""
        while(text != "exit\n"):
            text = input("$:")
            # self.serialPort.write(' ' + text + '\n')
            msg = "KK4YWN: " + text + "\n"
            self.serialPort.write(msg.encode())

send = Sender(serialPort) 
receive = Receiver(serialPort) 
send.start() 
receive.start()
