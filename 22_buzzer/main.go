package main

import (
	"machine"

	"tinygo.org/x/drivers/tone"
)

func main() {
	bzrPin := machine.GPIO1
	pwm := machine.PWM0
	speaker, err := tone.New(pwm, bzrPin)
	if err != nil {
		panic(err)
	}

	gpioPins := []machine.Pin{machine.GPIO28, machine.GPIO29, machine.GPIO2}
	for c := range gpioPins {
		gpioPins[c].Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	}

	prev := make([]bool, len(gpioPins))
	for i := range prev {
		prev[i] = !gpioPins[i].Get()
	}

	sound := []tone.Note{
		tone.C6,
		tone.E6,
		tone.G6,
	}

	for {
		for i := range gpioPins {
			current := gpioPins[i].Get()
			if prev[i] != current {
				prev[i] = current
				if !current {
					speaker.SetNote(sound[i])
				} else {
					speaker.Stop()
				}
			}
		}
	}
}
