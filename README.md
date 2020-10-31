# go-xinput

[![GoDoc](https://godoc.org/github.com/oxzi/go-xinput?status.svg)](https://godoc.org/github.com/oxzi/go-xinput) ![CI](https://github.com/oxzi/go-xinput/workflows/CI/badge.svg)

Small [Go][golang] library for limited interaction with the [X Input Device Extension Library][xorg-xinput] (`<X11/extensions/XInput.h>`).
This allows listing input devices and listening for events, e.g., pressing keys.


## Examples

### xinput-list

Reimplementation of [`xinput list`][xinput].

```
$ go run examples/xinput-list/main.go
Virtual core pointer                            id=2    [master pointer]
Virtual core keyboard                           id=3    [master keyboard]
Virtual core XTEST pointer                      id=4    [slave pointer]
Virtual core XTEST keyboard                     id=5    [slave keyboard]
Power Button                                    id=6    [slave keyboard]
Video Bus                                       id=7    [slave keyboard]
Sleep Button                                    id=8    [slave keyboard]
Integrated Camera: Integrated C                 id=10   [slave keyboard]
AT Translated Set 2 keyboard                    id=11   [slave keyboard]
TPPS/2 IBM TrackPoint                           id=13   [slave pointer]
ThinkPad Extra Buttons                          id=14   [slave keyboard]
FiiO DigiHug USB Audio                          id=9    [slave keyboard]
SynPS/2 Synaptics TouchPad                      id=12   [slave pointer]
```


### keylogger

Show key, button, or motion events from an input device.
Compared to other Go-based Linux _keyloggers_ this does not require extended permissions.

```
# Keyboard from the list above
$ go run examples/keylogger/main.go 11
{key press 37 map[]}
{key release 37 map[]}
{key press 37 map[]}
{key press 56 map[]}
{key release 37 map[]}
{key release 56 map[]}
{key press 45 map[]}

# Mouse from the list above
$ go run examples/keylogger/main.go 13
{motion 0 map[0:2299 1:866]}
{motion 0 map[0:2300 1:865]}
{motion 0 map[0:2300 1:864]}
{motion 0 map[0:2304 1:861]}
{motion 0 map[0:2305 1:859]}
{motion 0 map[1:857]}
{motion 0 map[1:856]}
{button press 1 map[]}
{button release 1 map[]}
{motion 0 map[1:856]}
```


[golang]: https://golang.org/
[xinput]: https://gitlab.freedesktop.org/xorg/app/xinput/
[xorg-xinput]: https://www.x.org/releases/X11R7.7/doc/libXi/inputlib.html
