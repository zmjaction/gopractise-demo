package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	Verbose bool
	Source  string
	Profile string
	// Port int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "newApp",
	Short: "This is a short description of the newApp tool",
	Long:  `This is a longer description of the newApp tool. It can be multiple lines.`,
	// 在PreRun函数之前执行
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PersistentFlags is running")
	},
	// 在Run函数之前执行
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PreRun is running")
	},
	// 执行命令时调用的函数
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("newApp is a CLI tool for creating new applications")
		fmt.Printf("rootCmd Verbose: %v\n", Verbose)
		fmt.Printf("rootCmdSource: %v\n", Source)
		// 从 viper 获取配置值
		fmt.Printf("config: %v\n", viper.AllSettings())
		fmt.Printf("port: %v\n", viper.GetInt("server.port"))
		fmt.Printf("level: %v\n", viper.GetString("log.level"))

	},
	// 在Run函数之后执行
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PostRun is running")
	},
	// 在PostRun函数之后执行
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PersistentPostRun is running")
	},
	// TraverseChildren 设为 true 时，在执行命令时会遍历所有子命令。
	// 这意味着在处理命令时，会递归地检查并处理当前命令的所有子命令，
	// 常用于需要对整个命令树进行操作的场景，例如查找特定的命令或者应用全局设置等。
	TraverseChildren: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	// 初始化配置文件
	// cobra.OnInitialize(initConfig)
	// 定义 --config 标志，用于指定配置文件的路径
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	// 通过 viper 自动绑定命令行标志与配置文件，./newApp --port 9098
	// rootCmd.PersistentFlags().Int("port", 9090, "服务器端口")
	// rootCmd.PersistentFlags().String("level", "info", "日志级别")
	// 持久标识
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	// 本地标识
	rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
	// 必选参数
	// rootCmd.Flags().StringVarP(&Profile, "profile", "p", "", "AWS profile (required)")
	// rootCmd.MarkFlagRequired("profile")
	// 通过 viper 自动绑定命令行标志与配置文件，./newApp --port 9098
	// viper.BindPFlag("server.port", rootCmd.PersistentFlags().Lookup("port"))
	// viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("level"))

}

// initConfig 读取并解析配置文件
func initConfig() {
	if cfgFile != "" {
		// 使用指定的配置文件
		viper.SetConfigFile(cfgFile)
	} else {
		// 默认查找配置文件
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)   // 查找 ~/.app.yaml
		viper.AddConfigPath(".")    // 查找当前目录
		viper.SetConfigType("yaml") // 设置默认类型
		viper.SetConfigName(".app") // 配置文件名 (.app.yaml)
	}
	// 设置默认值
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("log.level", "info")

	// 读取环境变量
	viper.AutomaticEnv()
	// 环境变量前缀：APP_SERVER_PORT
	viper.SetEnvPrefix("APP")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
	}
}
