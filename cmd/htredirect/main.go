package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/lugosieben/htredirect/config"
	"github.com/lugosieben/htredirect/internal/redirectserver"
)

func main() {
	fmt.Printf("htredirect %s\n", config.VERSION)
	config.Load()

	go redirectserver.Start(config.Port)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	fmt.Println("exiting")
}
