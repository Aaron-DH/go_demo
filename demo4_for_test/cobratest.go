package main

import (
	"github.com/spf13/cobra"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	command := app.NewCommand()

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
			for {

			}
		},
	}
	return cmd
}
