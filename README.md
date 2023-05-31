# dstarchat
This allows keyboard to keyboard chat over dv using an ic-9700. It's hakish and for testing only.

You need a LINUX pc connected to an ic-9700 via the usb port.

Open the ic-9700 advanced manual to page 11-23 and follow the instructions for enabling data in the radio.
Enable fast data as described on page 11-24.

Download gui.go into a folder
cd to said folder.
DO: go mod init gui.go
DO: go run gui.go

You do not have to specify a channel name. If you don't, you will be joined to 'Lobby'. This has no meaning so don't take it seriously.
Do enter your callsign so others will know who you are.

This will work with several chatters.

Have fun and let me know if you mod it. I have lots of ideas but I write code very slowly.
