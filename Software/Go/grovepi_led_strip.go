package main

import (
	"./grovepi"
	"fmt"
	"time"
)

func main() {
	var g grovepi.GrovePi
	g = *grovepi.InitGrovePi(0x04)
	err := g.PinMode(grovepi.D2, "output")
	if err != nil {
		fmt.Println(err)
	}
	g.ResetLedStrip()
	for {
		//g.DigitalWrite(grovepi.D2, 1)
		//time.Sleep(500 * time.Millisecond)
		//g.DigitalWrite(grovepi.D2, 0)
		fmt.Println("Start")
		start := time.Now()
		err := g.SetLedStripColor(0, 0, 255)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(time.Now().Sub(start))
		time.Sleep(500 * time.Millisecond)
		err = g.SetLedStripColor(255, 0, 0)
                if err != nil {
                        fmt.Println(err.Error())
			return
                }
		time.Sleep(500 * time.Millisecond)
		err = g.SetLedStripColor(0, 255, 0)
                if err != nil {
                        fmt.Println(err.Error())
			return
                }
		time.Sleep(500 * time.Millisecond)
		//return
	}
}
