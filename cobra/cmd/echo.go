package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
)

var echoCmd = &cobra.Command{
	Use:   "echo [参数1] [参数2]",
	Short: "回显两个参数",
	Args:  cobra.ExactArgs(2), // 必须恰好2个参数
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("参数1: %s\n", args[0])
		fmt.Printf("参数2: %s\n", args[1])
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)
}
