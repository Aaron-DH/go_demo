package main

import (
	"math/rand"
	"os"
	"time"

	"demo6_k8s_learn/app"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	command := app.NewKubeletCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
