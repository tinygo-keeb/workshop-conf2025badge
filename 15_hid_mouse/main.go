package main

import (
	"machine"
	"machine/usb/hid/mouse"
	"time"
)

func main() {
	btn := machine.GPIO29
	btn.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	m := mouse.Port()
	for {
		if !btn.Get() {
			m.Press(mouse.Left)
		} else {
			m.Release(mouse.Left)
		}
		time.Sleep(1 * time.Millisecond)
	}
}
