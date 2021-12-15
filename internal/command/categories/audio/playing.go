package audio

import (
	"strings"

	"github.com/Pauloo27/aryzona/internal/command"
	"github.com/Pauloo27/aryzona/internal/discord/voicer"
	"github.com/Pauloo27/aryzona/internal/utils"
)

var PlayingCommand = command.Command{
	Name: "playing", Aliases: []string{"np", "nowplaying", "tocando"},
	Description: "Show what is playing now",
	Handler: func(ctx *command.CommandContext) {
		vc := voicer.GetExistingVoicerForGuild(ctx.GuildID)
		if vc == nil {
			ctx.Error("Bot is not connect to a voice channel")
			return
		}
		playable := vc.Playing()

		if playable == nil {
			ctx.Error("Nothing playing...")
			return
		}

		embed := buildPlayableInfoEmbed(playable, vc).WithTitle("Now playing: " + playable.GetName())

		if vc.Queue.Size() > 1 {
			sb := strings.Builder{}
			next := vc.Queue.All()[1:]
			limit := len(next)
			if len(next) > 10 {
				limit = 10
			}
			for _, item := range next[:limit] {
				title, artist := item.GetFullTitle()
				if artist == "" {
					sb.WriteString(utils.Fmt("  -> %s\n", title))
				} else {
					sb.WriteString(utils.Fmt("  -> %s - %s\n", artist, title))
				}
			}
			if len(next) > 10 {
				sb.WriteString("_... and more ..._")
			}
			embed.WithField("**Coming next:**", sb.String())
		}

		ctx.SuccessEmbed(embed)
	},
}