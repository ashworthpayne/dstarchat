package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/tarm/serial"
	// for later "github.com/spf13/viper"
)

// Set some vars- need to move these to a config file
var serialDevice = "/dev/ttyUSB1"
var c = &serial.Config{Name: serialDevice, Baud: 9600}
var s, err = serial.OpenPort(c)

func readFromRadio(s *serial.Port, channel chan []byte) {
	for {
		buf := make([]byte, 1024)
		bytesRead, _ := s.Read(buf)
		//fmt.Println(string(buf[:bytesRead]))
		channel <- buf[:bytesRead]
	}
}

func writeToRadio(s *serial.Port, channel chan []byte) {
	for {
		data := <-channel
		//fmt.Println(string(data))
		s.Write(data)
		s.Flush()
	}
}

func drawchat(channel string, username string) {
	// Create a new GUI.
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
		return
	}
	defer g.Close()
	g.Cursor = true

	// Update the views when terminal changes size.
	g.SetManagerFunc(func(g *gocui.Gui) error {
		termwidth, termheight := g.Size()
		_, err := g.SetView("output", 0, 0, termwidth-1, termheight-4)
		if err != nil {
			return err
		}
		_, err = g.SetView("input", 0, termheight-3, termwidth-1, termheight-1)
		if err != nil {
			return err
		}
		return nil
	})

	// Terminal width and height.
	termwidth, termheight := g.Size()

	// Output.
	ov, err := g.SetView("output", 0, 0, termwidth-1, termheight-4)
	if err != nil && err != gocui.ErrUnknownView {
		log.Println("Failed to create output view:", err)
		return
	}
	ov.Title = " Messages  -  <" + channel + "> "
	ov.FgColor = gocui.ColorRed
	ov.Autoscroll = true
	ov.Wrap = true

	// Send a welcome message.
	_, err = fmt.Fprintln(ov, "<Net Control> OpenRRC v0.1.0")
	if err != nil {
		log.Println("Failed to print into output view:", err)
	}
	_, err = fmt.Fprintln(ov, "<Net Control> Press Ctrl-C to quit")
	if err != nil {
		log.Println("Failed to print into output view:", err)
	}

	// Input.
	iv, err := g.SetView("input", 0, termheight-3, termwidth-1, termheight-1)
	if err != nil && err != gocui.ErrUnknownView {
		log.Println("Failed to create input view:", err)
		return
	}
	iv.Title = " New Message  -  <" + username + "> "
	iv.FgColor = gocui.ColorWhite
	iv.Editable = true
	err = iv.SetCursor(0, 0)
	if err != nil {
		log.Println("Failed to set cursor:", err)
		return
	}

	// Bind Ctrl-C so the user can quit.
	err = g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	})
	if err != nil {
		log.Println("Could not set key binding:", err)
		return
	}

	// Bind enter key to input to send new messages.
	err = g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, iv *gocui.View) error {
		// Read buffer from the beginning.
		iv.Rewind()

		// Get output view and print.
		ov, err := g.View("output")
		if err != nil {
			log.Println("Cannot get output view:", err)
			return err
		}
		_, err = fmt.Fprintf(ov, "\u001b[31m<%s>: %s", username, iv.Buffer())
		if err != nil {
			log.Println("Cannot print to output view:", err)
		}
		radioChan := make(chan []byte)
		go writeToRadio(s, radioChan)

		// Send message if text was entered.
		if len(iv.Buffer()) >= 2 {

			radioChan <- []byte("<" + username + ">: " + "<" + channel + ">" + iv.Buffer())

			// Reset input.
			iv.Clear()

			// Reset cursor.
			err = iv.SetCursor(0, 0)
			if err != nil {
				log.Println("Failed to set cursor:", err)
			}
			return err
		}
		return nil
	})
	if err != nil {
		log.Println("Cannot bind the enter key:", err)
	}

	// Set the focus to input.
	_, err = g.SetCurrentView("input")
	if err != nil {
		log.Println("Cannot set focus to input view:", err)
	}

	successChannel := make(chan []byte)

	if err == nil {
		defer s.Close()
		go readFromRadio(s, successChannel)
	} else {
		log.Fatal("Failed to connect to the Radio")
	}
	go func() {
		for {
			response := <-successChannel
			ov, err := g.View("output")
			if err != nil {
				log.Println("Cannot get output view:", err)
				return
			}
			_, err = fmt.Fprintf(ov, "\u001b[32m%s", response)
			if err != nil {
				log.Println("Cannot print to output view:", err)
			}
			// Refresh view
			g.Update(func(g *gocui.Gui) error {
				return nil
			})
		}
	}()
	// Start the main loop.
	err = g.MainLoop()
	log.Println("Main loop has finished:", err)
}

func main() {
	// Get channel and username.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Channel Name: ")
	channel, err := reader.ReadString('\n')
	if channel == "\n" {
		channel = "Lobby\n"
	}
	if err != nil {
		log.Println("Could not set channel:", err)
	}
	fmt.Print("Enter your callsign: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Could not set callsign:", err)
	}

	// Create the GUI.
	drawchat(strings.TrimSuffix(channel, "\n"), strings.TrimSuffix(username, "\n"))
}
