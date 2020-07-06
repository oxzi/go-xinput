package xinput

// This file wraps xopendisplay(3) to allow usage of the display connection functions.

/*
#cgo LDFLAGS: -lX11

#include <X11/Xlib.h>
*/
import "C"

// XOpenDisplay returns a Display to be used in the functions. The displayName might be nil.
func XOpenDisplay(displayName *C.char) *C.Display {
	return C.XOpenDisplay(displayName)
}

// XCloseDisplay closes the connection established in XOpenDisplay.
func XCloseDisplay(display *C.Display) C.int {
	return C.XCloseDisplay(display)
}
