package main

import (
	"fmt"

	"github.com/mrverdant13/web-demo/go-app/config"
)

func main() {
	config := config.Get()
	fmt.Printf("%+v\n", config)
}
