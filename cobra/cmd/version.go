package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var VersionFlag string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of newApp",
	Long:  `All software has versions. This is newApp's`,
	Run: func(cmd *cobra.Command, args []string) {
		for i, arg := range args {
			fmt.Printf("参数 %d: %s\n", i+1, arg)
		}
		fmt.Println("newApp v0.1")
		fmt.Printf("version Verbose: %v\n", Verbose)
		fmt.Printf("version VersionFlag: %v\n", VersionFlag)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().StringVarP(&VersionFlag, "flag", "f", "false", "version flag")
}
