package actions

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bmc-toolbox/bmcfwupd/utils"
	"github.com/bmc-toolbox/bmclib/discover"
	"github.com/spf13/viper"
)

type firmware interface {
	Close() error
	HardwareType() string
	Model() (string, error)
	UpdateFirmware(string, string) (bool, error)
	Vendor() string
	Version() (string, error)
}

func connect(host string) (f firmware, err error) {
	username := viper.GetString("common.user")
	password := viper.GetString("common.pass")

	conn, err := discover.ScanAndConnect(host, username, password)
	if err != nil {
		return f, err
	}
	if f, ok := conn.(firmware); ok {
		return f, err
	}

	return f, fmt.Errorf("unable to detect device")
}

func CheckAndUpdate(hostFile string, update bool) (err error) {
	hosts := utils.ReadFile(hostFile)
	wg := sync.WaitGroup{}

	for _, host := range hosts {
		host := host
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			obj, err := connect(host)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer obj.Close()

			latest := false
			vendor := strings.ToLower(obj.Vendor())
			hardware := obj.HardwareType()
			runningVersion, _ := obj.Version()
			currentVersion := viper.GetString(fmt.Sprintf("%s.%s_version", vendor, hardware))

			if runningVersion == currentVersion {
				latest = true
			}

			if update {
				if latest {
					fmt.Printf("host=%s vendor=%s hardware=%sversion=%s latest=%t device_update=%t\n", host, vendor, hardware, runningVersion, latest, false)
				} else {
					currentFirmware := viper.GetString(fmt.Sprintf("%s.%s_firmware", vendor, hardware))
					source := viper.GetString(fmt.Sprintf("%s.source", vendor))
					result, err := obj.UpdateFirmware(source, currentFirmware)
					if err != nil {
						fmt.Printf("host=%s vendor=%s hardware=%s version=%s latest=%t device_update=%t err=%s\n", host, vendor, hardware, runningVersion, latest, result, err)
						return
					}
				}
			} else {
				fmt.Printf("host=%s vendor=%s hardware=%s version=%s latest=%t\n", host, vendor, hardware, runningVersion, latest)
			}
		}(&wg)
	}
	wg.Wait()

	return err
}
