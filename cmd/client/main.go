package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	address := flag.String("address", "localhost:3223", "Address of the database server")
	timeout := flag.Duration("timeout", 5*time.Second, "Timeout for connection and operations")
	flag.Parse()

	conn, err := net.DialTimeout("tcp", *address, *timeout)
	if err != nil {
		fmt.Printf("Failed to connect to server: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to WalGoDb server. Type 'exit' to quit or 'help' for help.")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("WalGoDb> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			break
		}
		if strings.ToLower(input) == "help" {
			printHelp()
			continue
		}

		_, err := fmt.Fprintf(conn, "%s\n", input)
		if err != nil {
			fmt.Printf("Error sending to server: %v\n", err)
			continue
		}

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading from server: %v\n", err)
			continue
		}

		fmt.Print(response)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
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
