package main

import (
	"machine"
	"time"
)

func main() {
	machine.GPIO29.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	machine.GPIO28.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	for {
		if !machine.GPIO29.Get() {
			println("上側のボタンが押されました")
		} else {
		}

		if !machine.GPIO28.Get() {
			println("下側のボタンが押されました")
		} else {
		}

		time.Sleep(100 * time.Millisecond)
	}
}
