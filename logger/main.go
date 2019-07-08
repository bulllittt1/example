package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// var logger *zap.Logger

func main() {
	rawJSON := []byte(`{
		  "level": "info",
		  "encoding": "json",
		  "outputPaths": ["stdout", "current"],
		  "errorOutputPaths": ["stderr"],
		  "initialFields": {"foo": "bar"},
		  "encoderConfig": {
		    "messageKey": "message",
		    "levelKey": "level",
		    "levelEncoder": "lowercase"
		  }
		}`)

	cfg := zap.NewProductionConfig()
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Debug("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "http://google.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	defer func() {
		recover()
		log.Println("panic recovered")
		os.Exit(0)
	}()

	log.Println("start")

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-time.After(3 * time.Second):
			log.Println("finished")
			os.Exit(0)
		case <-sigCh:
			log.Println("panic")
			panic(nil)
		}
	}
}
