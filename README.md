# dstarchat
This allows keyboard to keyboard chat over dv using an ic-9700. It's hakish and for testing only.

You need a LINUX pc connected to an ic-9700 via the usb port.

Open the ic-9700 advanced manual to page 11-23 and follow the instructions for enabling data in the radio. Please ignore the cabling diagram. The manual is trying to depict both cabling options but only one is required. Most will use the USB port. Once finished, proceed to enable fast data as described on page 11-24.

Its not clearly stated, but fast data mode adds a data subcarrier for the digital voice mode. Its essentially a 9600 baud serial link. I've heard people complain about being unable to run a TNC at 9600 baud on the 9700. Presumably, Icom feels fast data mode is superior to FM + TNC.

What is interersting about DV Fast Data mode is you can pick up the mic and talk at any time. And if a packet is ready, it will stay in queue until you are finished talking. 

Download gui.go into a folder
cd to said folder.
DO: go mod init gui.go
DO: go run gui.go

You do not have to specify a channel name. If you don't, you will be joined to 'Lobby'. This has no meaning so don't take it seriously.
Do enter your callsign so others will know who you are.

This will work with several chatters.

Have fun and let me know if you mod it. I have lots of ideas but I write code very slowly.
