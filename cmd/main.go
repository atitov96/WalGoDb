package main

import (
	"atitov96/walgodb/internal/compute"
	"atitov96/walgodb/internal/storage"
	"atitov96/walgodb/logger"
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"os"
)

func main() {
	log, err := logger.NewLogger()
	if err != nil {
		fmt.Println("failed to initialize logger")
	}
	defer func(log *zap.Logger) {
		err := log.Sync()
		if err != nil {
			fmt.Println("failed to sync logger")
		}
	}(log)

	parser := compute.NewParser()
	engine := storage.NewInMemoryEngine()
	computeLayer := compute.NewComputeLayer(parser, engine, log)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the expression to evaluate:")
	for {
		fmt.Print(">> ")
		expression, _ := reader.ReadString('\n')
		expression = expression[:len(expression)-1]
		if expression == "exit" {
			break
		}
		result, err := computeLayer.Execute(expression)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}
}
