package main

import (
	"fmt"
	"os"
	"signal"
	"syscall"
	"time"
)

func main() {
	shutdown := make(chan struct{})

	go func() {
		select {
		case c := <-shutdown:
			fmt.Println("shutdown", c)
			return
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT)

	s := <-c
	close(shutdown)
	fmt.Println("Got signal:", s)
	time.Sleep(100)
}
