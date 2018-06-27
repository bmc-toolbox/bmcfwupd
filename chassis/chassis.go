package chassis

import (
	"fmt"
	"strings"

	"github.com/bmc-toolbox/bmcfwupd/utils"
	"github.com/bmc-toolbox/bmclib/devices"
	"github.com/bmc-toolbox/bmclib/discover"
	"github.com/spf13/viper"
)

func connect(host string) (chassis devices.BmcChassis, err error) {
	username := viper.GetString("common.chassis_user")
	password := viper.GetString("common.chassis_pass")

	conn, err := discover.ScanAndConnect(host, username, password)
	if err != nil {
		return chassis, err
	}
	if chassis, ok := conn.(devices.BmcChassis); ok {
		return chassis, err
	}

	return chassis, err
}

func CheckFirmware(chassisfile string) (err error) {
	c7000CurrentVersion := viper.GetString("hp.c7000_version")
	m1000eCurrentVersion := viper.GetString("dell.m1000e_version")

	hosts := utils.ReadFile(chassisfile)
	for _, host := range hosts {
		obj, err := connect(host)
		if err != nil {
			return err
		}

		latest := false
		vendor := strings.ToLower(obj.Vendor())
		runningVersion, _ := obj.GetFirmwareVersion()

		switch vendor {
		case "dell":
			if runningVersion == m1000eCurrentVersion {
				latest = true
			}
		case "hp":
			if runningVersion == c7000CurrentVersion {
				latest = true
			}
		}
		fmt.Printf("chassis=%s vendor=%s latest=%t\n", host, vendor, latest)
	}

	return err
}
