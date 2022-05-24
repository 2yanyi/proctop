package internal

import (
	"github.com/jaypipes/ghw"
	"strings"
)

func graphics() []string {
	drivers := make([]string, 0)
	info, err := ghw.GPU()
	if err != nil {
		Stderr(err.Error())
	}
	if info == nil {
		return nil
	}
	for _, driver := range info.GraphicsCards {
		if driver == nil {
			continue
		}
		if driver.DeviceInfo == nil {
			continue
		}
		if driver.DeviceInfo.Product == nil {
			continue
		}
		if driver.DeviceInfo.Product.Name == "SVGA II Adapter" {
			continue
		}
		if strings.Contains(driver.DeviceInfo.Product.Name, "Graphics") {
			continue
		}
		drivers = append(drivers, driver.DeviceInfo.Product.Name)
	}
	return drivers
}
