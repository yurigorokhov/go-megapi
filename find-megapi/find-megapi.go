package main

import (
	"fmt"
	"github.com/yurigorokhov/go-megapi"
)

func main() {
	dev, err := megapi.Find_megapi_usb_device()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Found device: {}", dev)
	}
}
