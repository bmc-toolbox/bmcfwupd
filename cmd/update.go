package cmd

import (
	"fmt"
	"os"

	"github.com/bmc-toolbox/bmcfwupd/actions"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: `update firmware`,
	Long:  `update firmware`,
	Run: func(cmd *cobra.Command, args []string) {
		err := actions.CheckAndUpdate(hostFile, true)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&hostFile, "file", "f", "", "file containing hostname/ip")
	updateCmd.MarkFlagRequired("file")
}
