package main

import (
	"context"
	"fmt"
	"jetsend_opens/lib"
	"jetsend_opens/shared/helpers"
	"jetsend_opens/shared/log"
	"net/http"
	"os"
)

func main() {
	// Context to cancel functions
	ctx, cancel := context.WithCancel(context.Background())

	if len(os.Args) != 3 {
		fmt.Println("Usage: main  [path to config directory] [path to log directory]:", os.Args)
		return
	}

	configPath := os.Args[1]
	logPath := os.Args[2]

	lib.ParseConfiguration(configPath, logPath)
	lib.InitializeKafa(ctx)

	router := log.Logger(lib.GetRoutes())

	go helpers.CaptureSignal(cancel)

	fmt.Println("Starting opens HTTP Server ", lib.ServiceConfig.Service.Port)
	log.Debug(ctx, "Starting opens HTTP Server ", lib.ServiceConfig.Service.Port)

	err := http.ListenAndServe(":"+lib.ServiceConfig.Service.Port, router)
	log.Error(ctx, "opens HTTP ListenAndServe err : ", err)
}
