package permissions

import (
	"os"

	"github.com/Pauloo27/aryzona/internal/command"
)

var MustBeOwner = command.CommandPermission{
	Name: "be the bot owner",
	Checker: func(ctx *command.CommandContext) bool {
		return ctx.AuthorID == os.Getenv("DC_BOT_OWNER_ID")
	},
}
