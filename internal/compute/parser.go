package compute

import (
	"errors"
	"regexp"
	"strings"
)

type parser struct{}

func NewParser() Parser {
	return &parser{}
}

func (p *parser) Parse(expression string) (Command, error) {
	parts := strings.Fields(expression)
	if len(parts) == 0 {
		return Command{}, errors.New("empty expression")
	}
	commandType := strings.ToUpper(parts[0])
	args := parts[1:]

	if !isValidCommandType(commandType) {
		return Command{}, errors.New("unknown command")
	}

	if !areValidArguments(args) {
		return Command{}, errors.New("invalid arguments: must be alphanumeric")
	}

	return Command{Type: commandType, Args: args}, nil
}

func isValidCommandType(commandType string) bool {
	return commandType == "SET" || commandType == "GET" || commandType == "DEL"
}

func areValidArguments(args []string) bool {
	for _, arg := range args {
		if !isAlphanumeric(arg) {
			return false
		}
	}
	return true
}

func isAlphanumeric(s string) bool {
	return regexp.MustCompile("^[a-zA-Z0-9_\\-]*$").MatchString(s)
}
