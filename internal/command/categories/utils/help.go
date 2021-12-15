package utils

import (
	"os"
	"strings"

	"github.com/Pauloo27/aryzona/internal/command"
	"github.com/Pauloo27/aryzona/internal/utils"
)

var HelpCommand = command.Command{
	Name: "help", Description: "List all commands",
	Aliases: []string{"h"},
	Handler: func(ctx *command.CommandContext) {
		sb := strings.Builder{}
		sb.WriteString("I'm a open source bot, here's my code: ")
		sb.WriteString(utils.Fmt("%s\n", os.Getenv("DC_BOT_REMOTE_REPO")))
		sb.WriteString("List of commands:\n")
		for _, cmd := range command.GetCommandList() {
			var permission string
			if cmd.Permission != nil {
				permission = utils.Fmt("(_requires you to... %s_)", cmd.Permission.Name)
			}
			var aliases string
			if len(cmd.Aliases) > 0 {
				aliases = utils.Fmt("(aka %s)", strings.Join(cmd.Aliases, ", "))
			}
			sb.WriteString(utils.Fmt(
				" - %s `%s%s` %s: **%s** %s\n", cmd.GetCategory().Emoji,
				command.Prefix, cmd.Name, aliases, cmd.Description, permission,
			))
		}
		ctx.Success(sb.String())
	},
}