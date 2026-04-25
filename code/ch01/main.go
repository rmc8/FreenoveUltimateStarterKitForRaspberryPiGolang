package main

import (
	"fmt"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

func main() {
	// ホスト（Raspi）を初期化
	if _, err := host.Init(); err != nil {
		panic(err)
	}

	// GPIO 17 (BCM 17) を取得
	p := gpioreg.ByName("GPIO17")
	if p == nil {
		fmt.Println("Failed to find GPIO17")
		return
	}

	fmt.Println("Blinking LED on GPIO17 (Physical Pin 11)...")

	for {
		// Low/High を切り替え
		p.Out(gpio.Low)
		fmt.Println("LED OFF")
		time.Sleep(time.Second)

		p.Out(gpio.High)
		fmt.Println("LED ON")
		time.Sleep(time.Second)
	}
}
