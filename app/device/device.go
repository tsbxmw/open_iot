package main

import (
	"fmt"
	common "github.com/tsbxmw/gin_common"
	"open_iot/device/transport/http"
	"os"
)

func main() {
	httpServer := http.HttpServer{}
	config := common.ServiceConfigImpl{}
	app, err := common.App("device", "device management for iot", httpServer, config)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
