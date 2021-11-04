package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"com-request-api/internal/app/retranslator"
)

func main() {

	sigs := make(chan os.Signal, 1)

	cfg := retranslator.Config{
		ChannelSize:    512,
		ConsumerCount:  20,
		ConsumeTimeout: 1 * time.Second,
		ConsumeSize:    10,
		ProducerCount:  28,
		WorkerCount:    2,
	}

	retranslator := retranslator.NewRetranslator(cfg)
	retranslator.Start()

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	retranslator.Close()
}
