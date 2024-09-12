package compute

import (
	"atitov96/walgodb/internal/storage"
	"errors"
	"go.uber.org/zap"
)

type Parser interface {
	Parse(expression string) (string, []string, error)
}

type Compute interface {
	Execute(expression string) (string, error)
}

type computeLayer struct {
	parser  Parser
	storage storage.Storage
	log     *zap.Logger
}

func NewComputeLayer(p Parser, s storage.Storage, l *zap.Logger) Compute {
	return &computeLayer{
		parser:  p,
		storage: s,
		log:     l,
	}
}

func (c *computeLayer) Execute(expression string) (string, error) {
	command, args, err := c.parser.Parse(expression)
	if err != nil {
		c.log.Error("failed to parse expression", zap.Error(err))
		return "", err
	}

	c.log.Info("executing command", zap.String("command", command), zap.Strings("args", args))

	switch command {
	case "SET":
		if len(args) != 2 {
			err := errors.New("SET requires 2 arguments")
			c.log.Error("failed to execute command", zap.Error(err))
			return "", err
		}
		c.storage.Set(args[0], args[1])
		return "OK", err
	case "GET":
		if len(args) != 1 {
			err := errors.New("GET requires 1 argument")
			c.log.Error("failed to execute command", zap.Error(err))
			return "", err
		}
		val, ok := c.storage.Get(args[0])
		if !ok {
			return "NOT FOUND", nil
		}
		return val, nil
	case "DEL":
		if len(args) != 1 {
			err := errors.New("DEL requires 1 argument")
			c.log.Error("failed to execute command", zap.Error(err))
			return "", err
		}
		c.storage.Delete(args[0])
		return "OK", nil
	default:
		return "", errors.New("unknown command")
	}
}
