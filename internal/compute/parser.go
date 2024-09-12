package compute

import (
	"errors"
	"strings"
)

type parser struct{}

func NewParser() Parser {
	return &parser{}
}

func (p *parser) Parse(expression string) (string, []string, error) {
	parts := strings.Fields(expression)
	if len(parts) == 0 {
		return "", nil, errors.New("empty expression")
	}
	command := parts[0]
	args := parts[1:]

	if command != "SET" && command != "GET" && command != "DEL" {
		return "", nil, errors.New("unknown command")
	}

	return command, args, nil
}
