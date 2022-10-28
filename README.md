# dstarchat
This allows keyboard to keyboard chat over dv using a ic-9700. It's quick and dirty.

You need a pc connected to an ic-9700 via the usb port.

Open the ic-9700 advanced manual to page 11-23 and follow the instructions for enabling data in the radio.
Enable fast data as described on page 11-24

edit chat.py, find kk4ywn and replace if with your callsign. 

For even quicker/dirtier serial comms, try
cat /dev/ttyUSB1 & cat > /dev/ttyUSB1
in a terminal.
type some txt and press enter.
maybe use this to get chat.py to a friend?
good luck!

