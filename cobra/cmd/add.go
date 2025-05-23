package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [数字1] [数字2]...",
	Short: "累加多个数字",
	Args:  cobra.MinimumNArgs(2), // 至少2个参数
	Run: func(cmd *cobra.Command, args []string) {
		// 累加逻辑...
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
