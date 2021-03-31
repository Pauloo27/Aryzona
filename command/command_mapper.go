package command

import "strings"

var commandMap = map[string]*Command{}

var Prefix string

func RegisterCommand(command *Command) {
	commandMap[strings.ToLower(command.Name)] = command
	for _, alias := range command.Aliases {
		commandMap[strings.ToLower(alias)] = command
	}
}
