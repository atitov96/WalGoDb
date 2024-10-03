package main

import (
	"atitov96/walgodb/internal/compute"
	"atitov96/walgodb/internal/storage"
	"atitov96/walgodb/pkg/logger"
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	log, err := logger.NewLogger()
	if err != nil {
		fmt.Println("failed to initialize logger")
		os.Exit(1)
	}

	parser := compute.NewParser()
	engine := storage.NewInMemoryEngine()
	computeLayer := compute.NewComputeLayer(parser, engine, log)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-shutdown
		log.Info("shutting down gracefully...")
		defer func(log *zap.Logger) {
			err := log.Sync()
			if err != nil {
				fmt.Println("failed to flush logger")
			}
		}(log)
		os.Exit(0)
	}()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to WalGoDb. Type 'help' for usage information or 'exit' to quit.")

	for {
		fmt.Print("WalGoDb> ")
		expression, _ := reader.ReadString('\n')

		expression = strings.TrimSpace(expression)
		if expression == "" {
			continue
		}
		if expression == "help" {
			printHelp()
			continue
		}
		if expression == "exit" {
			break
		}

		result, err := computeLayer.Execute(expression)
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			continue
		}
		fmt.Println(result)
	}
}

func printHelp() {
	fmt.Println(`
Available commands:
SET <key> <value> : Set a key-value pair
GET <key>         : Get the value for a key
DEL <key>         : Delete a key-value pair
help              : Show this help message
exit              : Exit the program

Note: Keys and values must be alphanumeric (including underscores and hyphens)
	`)
}
