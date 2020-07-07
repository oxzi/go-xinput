package main

// keylogger lists key events on an input device.
// Thus, this is kind of a "keylogger" which can be executed without any special permissions.

import (
	"fmt"
	"os"
	"strconv"

	"github.com/geistesk/go-xinput"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage %s: [device-id]\n", os.Args[0])
		os.Exit(1)
	}

	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to parse device-id: %v\n", err)
		os.Exit(1)
	}

	display := xinput.XOpenDisplay(nil)
	defer xinput.XCloseDisplay(display)

	for _, device := range xinput.GetXDeviceInfos(display) {
		if device.Id != uint64(id) {
			continue
		}

		if eventMap, err := xinput.NewEventMap(display, device); err != nil {
			fmt.Printf("Failed to create EventMap: %v\n", err)
		} else {
			eventMap.Print()
		}
	}
}
