package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "清理资源",
	Args:  cobra.NoArgs, // 不允许任何参数
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("清理完成")
	},
}

//func init() {
//	rootCmd.AddCommand(cleanCmd)
//}
