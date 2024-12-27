package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "gobuild",
	Long: `gobuild可以将当前目录下的main.go编译成可执行文件并放到GOPATH的bin目录下`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gobuild可以将当前目录下的main.go编译成可执行文件并放到GOPATH的bin目录下")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
