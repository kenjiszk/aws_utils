package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "awsutils",
	Short: "This is aws server utility tool.",
	Long:  "This is aws server utility tool.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "1.0",
	Long:  `1.0`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-keisan v1.0")
	},
}
