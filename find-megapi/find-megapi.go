package main

import (
	"fmt"
	"github.com/yurigorokhov/go-megapi"
	"time"
)

func main() {
	dev, err := megapi.Find_megapi_usb_device()
	if err != nil {
		panic(err)
	}
	fmt.Println("Found device: {}", dev)
	megaPi, err := megapi.NewMegaPi(dev)
	if err != nil {
		panic(err)
	}
	defer megaPi.Close()
	time.Sleep(2 * time.Second)

	// stop both motors
	megaPi.DcMotorStop(1)
	megaPi.DcMotorStop(2)

	fmt.Println("Running motors on ports 1 and 2")
	speeds := []int16{50, 100, 200, 300}
	for _, speed := range speeds {
		fmt.Printf("Speed: %v\n", speed)
		megaPi.DcMotorRun(1, speed)
		megaPi.DcMotorRun(2, -speed)
		time.Sleep(3 * time.Second)
	}

	// stop both motors
	megaPi.DcMotorStop(1)
	megaPi.DcMotorStop(2)
}
