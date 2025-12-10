package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/gophers"
)

func main() {
	machine.I2C1.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
		SDA:       machine.GPIO6,
		SCL:       machine.GPIO7,
	})

	display := ssd1306.NewI2C(machine.I2C1)
	display.Configure(ssd1306.Config{
		Address: 0x3C,
		Width:   128,
		Height:  64,
	})
	display.ClearDisplay()
	time.Sleep(50 * time.Millisecond)

	white := color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}

	data := []byte("ABCEF")
	for {
		display.ClearBuffer()
		data[0], data[1], data[2], data[3], data[4] = data[1], data[2], data[3], data[4], data[0]
		tinyfont.WriteLine(display, &gophers.Regular32pt, 5, 45, string(data), white)
		display.Display()
		time.Sleep(200 * time.Millisecond)
	}
}
