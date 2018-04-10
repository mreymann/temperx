package main

import (
	"fmt"
	"log"
	"time"
	"github.com/zserge/hid"
)

func main() {
	hid_path := "413d:2107:0000:01"
	cmd_raw := []byte{0x01, 0x80, 0x33, 0x01, 0x00, 0x00, 0x00, 0x00}

	hid.UsbWalk(func(device hid.Device) {
		info := device.Info()
		id := fmt.Sprintf("%04x:%04x:%04x:%02x", info.Vendor, info.Product, info.Revision, info.Interface)
		if id != hid_path {
			return
		}

		if err := device.Open(); err != nil {
			log.Println("Open error: ", err)
			return
		}

		defer device.Close()

		if _, err := device.Write(cmd_raw, 1*time.Second); err != nil {
			log.Println("Output report write failed:", err)
			return
		}

		if buf, err := device.Read(-1, 1*time.Second); err == nil {
			tmp := (float64(buf[2]) * 256 + float64(buf[3])) / 100
			hum := (float64(buf[4]) * 256 + float64(buf[5])) / 100
			fmt.Printf("Temperature: %v, Humidity: %v\n", tmp, hum)
		}
	})
}
