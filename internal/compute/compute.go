package compute

import (
	"atitov96/walgodb/internal/storage"
	"errors"
	"go.uber.org/zap"
	"sync/atomic"
	"time"
)

type Command struct {
	Type string
	Args []string
}

type Parser interface {
	Parse(expression string) (Command, error)
}

type Metrics struct {
	TotalQueries   uint64
	SuccessQueries uint64
	FailedQueries  uint64
	AverageLatency time.Duration
}

type Compute interface {
	Execute(expression string) (string, error)
	GetMetrics() Metrics
}

type computeLayer struct {
	parser  Parser
	storage storage.Storage
	log     *zap.Logger
	metrics Metrics
}

func NewComputeLayer(p Parser, s storage.Storage, l *zap.Logger) Compute {
	return &computeLayer{
		parser:  p,
		storage: s,
		log:     l,
	}
}

func (c *computeLayer) Execute(expression string) (string, error) {
	start := time.Now()
	atomic.AddUint64(&c.metrics.TotalQueries, 1)

	command, err := c.parser.Parse(expression)
	if err != nil {
		c.log.Error("failed to parse expression", zap.Error(err))
		atomic.AddUint64(&c.metrics.FailedQueries, 1)
		return "", err
	}

	c.log.Info("executing command", zap.String("command type", command.Type), zap.Strings("command args", command.Args))

	var result string
	switch command.Type {
	case "SET":
		result, err = c.handleSet(command.Args)
	case "GET":
		result, err = c.handleGet(command.Args)
	case "DEL":
		result, err = c.handleDel(command.Args)
	default:
		err = errors.New("unknown command")
	}

	if err != nil {
		atomic.AddUint64(&c.metrics.FailedQueries, 1)
	} else {
		atomic.AddUint64(&c.metrics.SuccessQueries, 1)
	}

	duration := time.Since(start)
	atomic.StoreInt64((*int64)(&c.metrics.AverageLatency), int64((c.metrics.AverageLatency*time.Duration(c.metrics.TotalQueries-1)+duration)/time.Duration(c.metrics.TotalQueries)))

	return result, err
}

func (c *computeLayer) GetMetrics() Metrics {
	return c.metrics
}

func (c *computeLayer) handleSet(args []string) (string, error) {
	if len(args) != 2 {
		return "", errors.New("SET requires 2 arguments")
	}
	c.storage.Set(args[0], args[1])
	return "OK", nil
}

func (c *computeLayer) handleGet(args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New("GET requires 1 argument")
	}
	val, ok := c.storage.Get(args[0])
	if !ok {
		return "NOT FOUND", nil
	}
	return val, nil
}

func (c *computeLayer) handleDel(args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New("DEL requires 1 argument")
	}
	c.storage.Delete(args[0])
	return "OK", nil
}
