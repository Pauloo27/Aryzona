package utils

import (
	"github.com/Pauloo27/aryzona/internal/command"
	"github.com/Pauloo27/aryzona/internal/discord"
	"github.com/google/uuid"
)

var UUIDCommand = command.Command{
	Name: "uuid", Aliases: []string{"gid", "id", "uid", "guid"},
	Description: "Generate an UUID",
	Handler: func(ctx *command.CommandContext) {
		id := uuid.New()
		ctx.SuccessEmbed(
			discord.NewEmbed().
				WithTitle("UUID v4").
				WithDescription(id.String()),
		)
	},
}
