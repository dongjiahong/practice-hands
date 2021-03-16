package cmd

import (
	"github.com/spf13/cobra"
)

var echoTime int

func init() {
	// 添加--time 和 -t 参数
	versionCmd.Flags().IntVarP(&echoTime, "time", "t", 1, "times to echo the input")
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Gonne",
	Run: func(cmd *cobra.Command, args []string) {
		println("cmd version is 0.0.1, time: ", echoTime)
	},
}
