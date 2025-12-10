package main

import (
	"fmt"
	"machine"
	"time"

	pio "github.com/tinygo-org/pio/rp2-pio"
	"github.com/tinygo-org/pio/rp2-pio/piolib"
	"tinygo.org/x/drivers/tone"
)

type WS2812B struct {
	Pin machine.Pin
	ws  *piolib.WS2812B
}

func NewWS2812B(pin machine.Pin) *WS2812B {
	s, _ := pio.PIO0.ClaimStateMachine()
	ws, _ := piolib.NewWS2812B(s, pin)
	ws.EnableDMA(true)
	return &WS2812B{
		ws: ws,
	}
}

func (ws *WS2812B) WriteRaw(rawGRB []uint32) error {
	return ws.ws.WriteRaw(rawGRB)
}

func main() {
	bzrPin := machine.GPIO1
	pwm := machine.PWM0
	speaker, err := tone.New(pwm, bzrPin)
	if err != nil {
		panic(err)
	}

	colors := []uint32{
		0xFFFFFFFF, 0xFFFFFFFF,
	}

	ws := NewWS2812B(machine.GPIO0)

	gpioPins := []machine.Pin{machine.GPIO29, machine.GPIO28}
	for c := range gpioPins {
		gpioPins[c].Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	}

	changed := true
	for {
		if !gpioPins[0].Get() {
			if colors[0] != 0x00000000 {
				fmt.Printf("sw1 pressed\n")
				colors[0] = 0x00000000
				speaker.SetNote(tone.B5)
				changed = true
			}
		} else {
			if colors[0] != 0xFFFFFFFF {
				colors[0] = 0xFFFFFFFF
				speaker.Stop()
				changed = true
			}
		}
		if !gpioPins[1].Get() {
			if colors[1] != 0x00000000 {
				fmt.Printf("sw2 pressed\n")
				colors[1] = 0x00000000
				speaker.SetNote(tone.D6)
				changed = true
			}
		} else {
			if colors[1] != 0xFFFFFFFF {
				colors[1] = 0xFFFFFFFF
				changed = true
				speaker.Stop()
			}
		}

		if changed {
			ws.WriteRaw(colors)
			time.Sleep(32 * time.Millisecond)
			changed = false
		}
	}
}
