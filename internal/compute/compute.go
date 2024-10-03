package compute

import (
	"atitov96/walgodb/internal/storage"
	"errors"
	"go.uber.org/zap"
)

type Command struct {
	Type string
	Args []string
}

type Parser interface {
	Parse(expression string) (Command, error)
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
	command, err := c.parser.Parse(expression)
	if err != nil {
		c.log.Error("failed to parse expression", zap.Error(err))
		return "", err
	}

	c.log.Info("executing command", zap.String("command type", command.Type), zap.Strings("command args", command.Args))

	switch command.Type {
	case "SET":
		return c.handleSet(command.Args)
	case "GET":
		return c.handleGet(command.Args)
	case "DEL":
		return c.handleDel(command.Args)
	default:
		return "", errors.New("unknown command")
	}
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
