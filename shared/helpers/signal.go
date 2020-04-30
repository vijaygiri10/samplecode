package helpers

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//CaptureSignal ...
func CaptureSignal(cancel context.CancelFunc) {
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		signalReceived := <-signalChan
		switch signalReceived {
		// kill -SIGINT XXXX or Ctrl+c
		case syscall.SIGINT:
			fmt.Printf("\ncaught sig: %+v\n", syscall.SIGINT)
			fmt.Printf("Gracefully shutting down server...\n")
			cancel()
			os.Exit(int(syscall.SIGKILL))
		// kill -SIGTERM XXXX
		case syscall.SIGTERM:
			fmt.Printf("\ncaught sig: %+v\n", syscall.SIGTERM)
			fmt.Printf("Gracefully shutting down server...\n")
			cancel()
			os.Exit(int(syscall.SIGTERM))
		default:
			fmt.Printf("\ncaught sig: %+v\n", signalReceived)
			fmt.Printf("Gracefully shutting down server...\n")
			cancel()
			os.Exit(1)
		}
	}
}
