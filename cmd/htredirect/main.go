package main

import (
	"fmt"

	"github.com/lugosieben/htredirect/config"
	"github.com/lugosieben/htredirect/web"
)

func main() {
	fmt.Printf("htredirect %s\n", config.VERSION)
	config.Load()
	web.Start(config.Port)
}
