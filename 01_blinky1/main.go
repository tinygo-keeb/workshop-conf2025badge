package main

import (
	"image/color"
	"machine"
	"time"

	pio "github.com/tinygo-org/pio/rp2-pio"
	"github.com/tinygo-org/pio/rp2-pio/piolib"
)

type WS2812B struct {
	Pin machine.Pin
	ws  *piolib.WS2812B
}

func NewWS2812B(pin machine.Pin) *WS2812B {
	s, _ := pio.PIO0.ClaimStateMachine()
	ws, _ := piolib.NewWS2812B(s, pin)
	return &WS2812B{
		ws: ws,
	}
}

func (ws *WS2812B) PutColor(c color.Color) {
	ws.ws.PutColor(c)
}

var (
	white = color.RGBA{R: 0x10, G: 0x10, B: 0x10, A: 0xFF}
	black = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}
)

func main() {
	// XIAO RP2040 のボード上の LED は電源制御が必要
	machine.NEO_PWR.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.NEO_PWR.High()

	ws := NewWS2812B(machine.NEOPIXEL)
	ws.PutColor(white)

	for {
		time.Sleep(time.Millisecond * 500)
		ws.PutColor(black)
		time.Sleep(time.Millisecond * 500)
		ws.PutColor(white)
	}
}
