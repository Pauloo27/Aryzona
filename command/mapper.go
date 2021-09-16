package command

import (
	"strings"

	"github.com/Pauloo27/logger"
)

var commandMap = map[string]*Command{}

var Prefix string

func RegisterCommand(command *Command) {
	if command.Name == "" {
		logger.Fatal("One command has no name")
	}
	if command.Description == "" {
		logger.Fatalf("Command %s has no description", command.Name)
	}
	commandMap[strings.ToLower(command.Name)] = command
	for _, alias := range command.Aliases {
		commandMap[strings.ToLower(alias)] = command
	}
}

// why a function? I think I did it that way, so the access to the
// command map was "harder" (the idea is to use RegisterCommand())
func GetCommandMap() map[string]*Command {
	return commandMap
}

func RegisterCategory(category CommandCategory) {
	if category.OnLoad != nil {
		category.OnLoad()
	}
	if category.Name == "" {
		logger.Fatal("One category has no name")
	}
	if category.Emoji == "" {
		logger.Fatalf("Category %s has no emoji", category.Name)
	}
	for _, cmd := range category.Commands {
		cmd.category = &category
		RegisterCommand(cmd)
	}
}
