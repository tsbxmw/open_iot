package main

import (
	"fmt"
	common "github.com/tsbxmw/gin_common"
	"open_iot/keng/transport/http"
	"os"
)

func main() {
	httpServer := http.HttpServer{}
	config := common.ServiceConfigImpl{}
	app, err := common.App("keng", "keng for open iot", httpServer, config)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
