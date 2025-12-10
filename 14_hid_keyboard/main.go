package main

import (
	"machine"
	"machine/usb/hid/keyboard"
	"time"
)

func main() {
	btn := machine.GPIO29
	btn.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	kb := keyboard.Port()
	for {
		if !btn.Get() {
			kb.Down(keyboard.KeyA)
		} else {
			kb.Up(keyboard.KeyA)
		}
		time.Sleep(1 * time.Millisecond)
	}
}
