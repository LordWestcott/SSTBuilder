package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	p := Project{}
	go p.Init(&wg)

	//exit if you get a SIGINT or SIGTERM or OS signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	go func() {
		sig := <-sigs
		println("Signal received: ", sig)
		println("Exiting...")
		close(sigs)
		os.Exit(0)
	}()

	wg.Wait()
	close(sigs)
}
