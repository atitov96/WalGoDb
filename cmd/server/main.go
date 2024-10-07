package main

import (
	"atitov96/walgodb/internal/compute"
	"atitov96/walgodb/internal/storage"
	"atitov96/walgodb/pkg/config"
	"atitov96/walgodb/pkg/logger"
	"context"
	"fmt"
	"go.uber.org/zap"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Server struct {
	listener       net.Listener
	computeLayer   compute.Compute
	log            *zap.Logger
	maxConnections int
	idleTimeout    time.Duration
	connections    sync.WaitGroup
	semaphore      chan struct{}
}

func (s *Server) Serve(ctx context.Context) {
	s.semaphore = make(chan struct{}, s.maxConnections)

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				return
			default:
				s.log.Error("Error accepting connections", zap.Error(err))
			}
		}

		select {
		case s.semaphore <- struct{}{}:
			s.connections.Add(1)
			go s.handleConnection(ctx, conn)
		default:
			s.log.Error("Max connections reached, closed new connection")
			conn.Close()
		}
	}
}

func (s *Server) Shutdown() {
	s.listener.Close()
	s.connections.Wait()
}

func (s *Server) handleConnection(ctx context.Context, conn net.Conn) {
	defer func() {
		conn.Close()
		<-s.semaphore
		s.connections.Done()
	}()

	for {
		err := conn.SetDeadline(time.Now().Add(s.idleTimeout))
		if err != nil {
			s.log.Error("Error setting connection deadline", zap.Error(err))
			return
		}

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				s.log.Info("Connection idle timeout")
			} else {
				s.log.Error("Error reading from connection", zap.Error(err))
			}
			return
		}

		query := string(buffer[:n])
		result, err := s.computeLayer.Execute(query)
		if err != nil {
			s.log.Error("Error executing query", zap.Error(err))
			_, writeErr := conn.Write([]byte(fmt.Sprintf("Error: %v\n", err)))
			if writeErr != nil {
				s.log.Error("Error writing error response", zap.Error(writeErr))
				return
			}
		} else {
			_, writeErr := conn.Write([]byte(fmt.Sprintf("%s\n", result)))
			if writeErr != nil {
				s.log.Error("Error writing response", zap.Error(writeErr))
				return
			}
		}

		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	log, err := logger.NewLogger(cfg.Logging.Level, cfg.Logging.Output)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer log.Sync()

	parser := compute.NewParser()
	engine := storage.NewInMemoryEngine()
	computeLayer := compute.NewComputeLayer(parser, engine, log)

	listener, err := net.Listen("tcp", cfg.Network.Address)
	if err != nil {
		log.Fatal("Failed to start TCP server", zap.Error(err))
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

	log.Info("TCP server started", zap.String("address", cfg.Network.Address))

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown
	log.Info("Shutting down gracefully...")
	cancel()
	server.Shutdown()
}
