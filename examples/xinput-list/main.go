package main

// xinput-list is a limited reimplementation of `xinput list`.

import (
	"fmt"

	"github.com/oxzi/go-xinput"
)

func main() {
	display := xinput.XOpenDisplay(nil)
	defer xinput.XCloseDisplay(display)

	for _, device := range xinput.GetXDeviceInfos(display) {
		fmt.Printf("%-40s\tid=%d\t[%v]\n", device.Name, device.Id, device.Use)
	}
}
