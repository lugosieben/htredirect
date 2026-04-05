package main

import (
	"fmt"

	"github.com/lugosieben/htredirect/config"
	"github.com/lugosieben/htredirect/internal/webserver"
)

func main() {
	fmt.Printf("htredirect %s\n", config.VERSION)
	config.Load()
	webserver.Start(config.Port)
}
