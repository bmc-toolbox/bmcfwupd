package cmd

import (
	"fmt"
	"os"

	"github.com/bmc-toolbox/bmcfwupd/chassis"
	"github.com/spf13/cobra"
)

var chassisfile string

func init() {
	chassisCmd.AddCommand(chassisCheckCmd)
	chassisCmd.AddCommand(chassisUpdateCmd)
	chassisCheckCmd.Flags().StringVarP(&chassisfile, "file", "f", "", "file containing hostname/ip")
	chassisCheckCmd.MarkFlagRequired("file")
	chassisUpdateCmd.Flags().StringVarP(&chassisfile, "file", "f", "", "file containing hostname/ip")
	chassisUpdateCmd.MarkFlagRequired("file")
}

var chassisCmd = &cobra.Command{
	Use:   "chassis",
	Short: `Chassis commands`,
	Long:  `Chassis commands`,
}

var chassisCheckCmd = &cobra.Command{
	Use:   "check",
	Short: `check firmware`,
	Long:  `check firmware`,
	Run: func(cmd *cobra.Command, args []string) {
		err := chassis.CheckFirmware(chassisfile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var chassisUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: `update firmware`,
	Long:  `update firmware`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updating chassis firmware")
	},
}
