package cmd

import (
	"fmt"
	"os"

	"github.com/bmc-toolbox/bmcfwupd/actions"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: `check firmware`,
	Long:  `check firmware`,
	Run: func(cmd *cobra.Command, args []string) {
		err := actions.CheckAndUpdate(hostFile, false)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().StringVarP(&hostFile, "file", "f", "", "file containing hostname/ip")
	checkCmd.MarkFlagRequired("file")
}
