package sobot

import (
	device "github.com/go-rod/rod/lib/devices"
)

type AccountInfo struct {
	Username, Password, FilePath, Pname, Caption string
}

var (
	rod_device = device.Device{
		Title:        "Laptop with touch",
		Capabilities: []string{},
		UserAgent:    "",
		Screen: device.Screen{
			DevicePixelRatio: 1,
			Horizontal: device.ScreenSize{
				Width:  800,
				Height: 760,
			},
			Vertical: device.ScreenSize{
				Width:  800,
				Height: 760,
			},
		},
	}
)
