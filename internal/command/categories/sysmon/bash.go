package sysmon

import (
	"os/exec"
	"strings"

	"github.com/Pauloo27/aryzona/internal/command"
	"github.com/Pauloo27/aryzona/internal/command/permissions"
	"github.com/Pauloo27/aryzona/internal/utils"
)

// should i remove it? probably...
/* #nosec G204 */
var Bash = command.Command{
	Name:        "bash",
	Description: "Eval a bash command",
	Permission:  &permissions.BeOwner,
	Arguments: []*command.CommandArgument{
		{
			Name: "command", Description: "command to execute", Required: true,
			RequiredMessage: "Missing command", Type: command.ArgumentText,
		},
	},
	Handler: func(ctx *command.CommandContext) {
		cmd := exec.Command("bash", "-c", (ctx.Args[0].(string)))
		buffer, err := cmd.CombinedOutput()
		output := string(buffer)
		output = strings.ReplaceAll(output, "`", "\\`")
		if err != nil {
			ctx.Error(utils.Fmt("Something went wrong:\n```\n%s\n```", output))
		} else {
			ctx.Success(utils.Fmt("Command ran successfully:\n```\n%s\n```", output))
		}
	},
}