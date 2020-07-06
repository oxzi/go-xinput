package xinput

// This file partially implements XDeviceInfo from xlistinputdevices(3).

/*
#cgo LDFLAGS: -lX11 -lXi

#include <X11/Xlib.h>
#include <X11/extensions/XInput.h>
*/
import "C"
import "unsafe"

// XDeviceInfoUse is a more Go-ish way to describe XDeviceInfo's use field.
// Here, the "master" and "slave" kinds for keyboards and pointers are implemented.
type XDeviceInfoUse string

const (
	KeyboardMaster XDeviceInfoUse = "master keyboard"
	KeyboardSlave  XDeviceInfoUse = "slave keyboard"
	PointerMaster  XDeviceInfoUse = "master pointer"
	PointerSlave   XDeviceInfoUse = "slave pointer"

	xDeviceInfoUse_Unknown XDeviceInfoUse = "unknown"
)

// parseUseField creates a XDeviceInfoUse from the use field.
func parseUseField(use C.int) XDeviceInfoUse {
	switch use {
	case C.IsXKeyboard:
		return KeyboardMaster

	case C.IsXExtensionKeyboard:
		return KeyboardSlave

	case C.IsXPointer:
		return PointerMaster

	case C.IsXExtensionPointer:
		return PointerSlave

	default:
		return xDeviceInfoUse_Unknown
	}
}

// XDeviceInfo is the partner data structure in Go with fewer fields.
type XDeviceInfo struct {
	Id   uint64
	Name string
	Use  XDeviceInfoUse
}

// GetXDeviceInfos fetches all XDeviceInfos from a Display.
func GetXDeviceInfos(display *C.Display) (devices []XDeviceInfo) {
	var numDevices C.int
	rawDevices := C.XListInputDevices(display, &numDevices)

	// Create a Go-ish slice of up to int32 C.XDeviceInfo, limited to numDevices.
	// Thanks to: https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices
	cDevices := (*[1 << 31]C.XDeviceInfo)(unsafe.Pointer(rawDevices))[:numDevices:numDevices]

	for _, cDevice := range cDevices {
		devices = append(devices, XDeviceInfo{
			Id:   uint64(cDevice.id),
			Name: C.GoString(cDevice.name),
			Use:  parseUseField(cDevice.use),
		})
	}

	C.XFreeDeviceList(rawDevices)

	return
}
