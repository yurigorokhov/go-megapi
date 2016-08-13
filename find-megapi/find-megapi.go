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
	time.Sleep(1 * time.Second)
	defer megaPi.Close()
	fmt.Println("Running motor on port 1 for 5 seconds")
	megaPi.MotorRun(1, 0)
	megaPi.MotorRun(2, -1)
	time.Sleep(1 * time.Second)
	megaPi.MotorRun(1, 50)
	megaPi.MotorRun(2, -50)
	time.Sleep(5 * time.Second)
	megaPi.MotorRun(1, 0)
	megaPi.MotorRun(2, -1)
}
