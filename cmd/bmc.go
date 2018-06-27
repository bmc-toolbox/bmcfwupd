package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var bmcfile string

func init() {
	bmcCmd.AddCommand(bmcCheckCmd)
	bmcCmd.AddCommand(bmcUpdateCmd)
	bmcCheckCmd.Flags().StringVarP(&bmcfile, "file", "f", "", "file containing hostname/ip")
	bmcCheckCmd.MarkFlagRequired("file")
	bmcUpdateCmd.Flags().StringVarP(&bmcfile, "file", "f", "", "file containing hostname/ip")
	bmcUpdateCmd.MarkFlagRequired("file")
}

var bmcCmd = &cobra.Command{
	Use:   "bmc",
	Short: `BMC commands`,
	Long:  `BMC commands`,
}

var bmcCheckCmd = &cobra.Command{
	Use:   "check",
	Short: `check firmware`,
	Long:  `check firmware`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("checking bmc firmware")
	},
}

var bmcUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: `update firmware`,
	Long:  `update firmware`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updating bmc firmware")
	},
}
