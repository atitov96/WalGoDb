package main

import (
	"atitov96/walgodb/internal/compute"
	"atitov96/walgodb/internal/storage"
	"atitov96/walgodb/pkg/config"
	"atitov96/walgodb/pkg/logger"
	"context"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	cfg := &config.Config{
		Network: config.NetworkConfig{
			Address:        "localhost:0",
			MaxConnections: 10,
			IdleTimeout:    time.Second * 5,
		},
	}

	log, _ := logger.NewLogger("info", "stdout")
	parser := compute.NewParser()
	engine := storage.NewInMemoryEngine()
	computeLayer := compute.NewComputeLayer(parser, engine, log)

	listener, err := net.Listen("tcp", cfg.Network.Address)
	if err != nil {
		t.Fatalf("Failed to start test server: %v", err)
	}

	server := &Server{
		listener:       listener,
		computeLayer:   computeLayer,
		log:            log,
		maxConnections: cfg.Network.MaxConnections,
		idleTimeout:    cfg.Network.IdleTimeout,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go server.Serve(ctx)

	// Test connection and basic operation
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatalf("Failed to connect to test server: %v", err)
	}

	_, err = fmt.Fprintf(conn, "SET test_key test_value\n")
	if err != nil {
		t.Fatalf("Failed to send SET command: %v", err)
	}

	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	if string(response[:n]) != "OK\n" {
		t.Errorf("Unexpected response for SET command: %s", string(response[:n]))
	}

	// Test connection limit
	connections := make([]net.Conn, cfg.Network.MaxConnections)
	for i := 0; i < cfg.Network.MaxConnections; i++ {
		connections[i], err = net.Dial("tcp", listener.Addr().String())
		if err != nil {
			t.Fatalf("Failed to open connection %d: %v", i, err)
		}
	}

	_, err = net.Dial("tcp", listener.Addr().String())
	if err == nil {
		t.Errorf("Expected connection to be refused, but it was accepted")
	}

	// Close connections
	for _, conn := range connections {
		conn.Close()
	}

	// Test idle timeout
	conn, _ = net.Dial("tcp", listener.Addr().String())
	time.Sleep(cfg.Network.IdleTimeout + time.Second)
	_, err = fmt.Fprintf(conn, "GET test_key\n")
	//if err == nil {
	//	t.Errorf("Expected error due to idle timeout, but got none")
	//}

	cancel()
	server.Shutdown()
}
