package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
	"time"
)

//声明变量用于接收命令行传入的参数值
var (
    name    string
)

func main() {
	rand.Seed(time.Now().UnixNano())

	command := NewCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Test Cobra command",
		Short: "Shot description",
		Long:  "Long description",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Test Run")
			fmt.Println(cmd.Flags())
			for {

			}
		},
	}

	fs := cmd.Flags()
	fs.StringVar(&name,  "name", "NonUser", "Your Name")
	return cmd
}
