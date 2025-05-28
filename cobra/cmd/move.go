package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var moveCmd = &cobra.Command{
	Use:   "move [源文件] [目标文件]",
	Short: "移动文件",
	Args:  cobra.ArbitraryArgs, // 允许任意数量的参数
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("警告：未指定要移动的文件")
			return
		}
		fmt.Printf("移动 %d 个文件\n", len(args))
	},
}

//func init() {
//	rootCmd.AddCommand(moveCmd)
//}
