package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var colorCmd = &cobra.Command{
	Use:       "color [颜色]",
	Short:     "选择颜色",
	Args:      cobra.OnlyValidArgs, // 参数必须在 ValidArgs 中
	ValidArgs: []string{"red", "green", "blue"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Printf("选择的颜色: %s\n", args[0])
		} else {
			fmt.Println("请选择一种颜色")
		}
	},
}

//func init() {
//	rootCmd.AddCommand(colorCmd)
//}
