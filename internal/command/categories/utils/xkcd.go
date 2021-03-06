package utils

import (
	"fmt"

	"github.com/Pauloo27/aryzona/internal/command"
	"github.com/Pauloo27/aryzona/internal/command/parameters"
	"github.com/Pauloo27/aryzona/internal/discord"
	"github.com/Pauloo27/aryzona/internal/providers/xkcd"
	"github.com/Pauloo27/aryzona/internal/utils"
)

var XkcdCommand = command.Command{
	Name: "xkcd", Description: "Get a xkcd comic",
	SubCommands: []*command.Command{
		&XkcdLatestSubCommand,
		&XkcdRandomSubCommand,
		&XkcdNumberSubCommand,
	},
}

var XkcdLatestSubCommand = command.Command{
	Name: "latest", Description: "Get the latest xkcd comic",
	Handler: func(ctx *command.CommandContext) {
		comic, err := xkcd.GetLatest()
		sendComic(ctx, comic, err)
	},
}

var XkcdRandomSubCommand = command.Command{
	Name: "random", Description: "Get a random xkcd comic",
	Handler: func(ctx *command.CommandContext) {
		comic, err := xkcd.GetRandom()
		sendComic(ctx, comic, err)
	},
}

var XkcdNumberSubCommand = command.Command{
	Name: "number", Description: "Get a xkcd comic by it's number",
	Aliases: []string{"num"},
	Parameters: []*command.CommandParameter{
		{
			Name: "number", Description: "number of the comic",
			Type: parameters.ParameterInt, Required: true,
		},
	},
	Handler: func(ctx *command.CommandContext) {
		comic, err := xkcd.GetByNum(ctx.Args[0].(int))
		sendComic(ctx, comic, err)
	},
}

func sendComic(ctx *command.CommandContext, comic *xkcd.Comic, err error) {
	if err != nil {
		ctx.Error("Cannot get comic =(")
		return
	}

	ctx.SuccessEmbed(
		discord.NewEmbed().
			WithTitle(fmt.Sprintf(
				"#%d - %s (%s/%s/%s)", comic.Num, comic.SafeTitle,
				comic.Year, utils.PadLeft(comic.Month, "0", 2), utils.PadLeft(comic.Day, "0", 2)),
			).
			WithURL(fmt.Sprintf("https://www.explainxkcd.com/wiki/index.php/%d", comic.Num)).
			WithImage(comic.Img).
			WithFooter(comic.Alt),
	)
}
