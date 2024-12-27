package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use: "name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("请输入名字")
			return
		}
		// 获取 GOPATH
		goPath, _ := exec.Command("go", "env", "GOPATH").Output()
		goPath = goPath[:len(goPath)-1]
		goPathStr := string(goPath) + "/bin/" + args[0]
		cmdstr := "go build -o " + goPathStr + " main.go"
		fmt.Println(cmdstr)
		buildCmd := exec.Command("go", "build", "-o", goPathStr, "main.go")
		if err := buildCmd.Run(); err != nil {
			fmt.Printf("Failed to build: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
