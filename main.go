package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Pauloo27/aryzona/command"
	"github.com/Pauloo27/aryzona/discord"
	"github.com/Pauloo27/aryzona/git"
	"github.com/Pauloo27/logger"
	"github.com/joho/godotenv"

	// import all command categories
	_ "github.com/Pauloo27/aryzona/command/categories/audio"
	_ "github.com/Pauloo27/aryzona/command/categories/sysmon"
	_ "github.com/Pauloo27/aryzona/command/categories/utils"
	"github.com/Pauloo27/aryzona/command/slash"
)

var commitHash, commitMessage string

func init() {
	logger.Info("Loading .env...")
	err := godotenv.Load()
	logger.HandleFatal(err, "Cannot load .env")
	logger.Success(".env loaded")

	git.CommitHash = commitHash
	git.CommitMessage = commitMessage
	git.RemoteRepo = os.Getenv("DC_BOT_REMOTE_REPO")
}

func main() {
	logger.Info("Connecting to Discord...")
	err := discord.Create(os.Getenv("DC_BOT_TOKEN"))
	if err != nil {
		logger.Fatal(err)
	}
	discord.AddDefaultListeners()
	err = discord.Connect()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Success("Connected to discord")

	command.Prefix = os.Getenv("DC_BOT_PREFIX")

	logger.Info("Updating slash commands, it may take a while...")
	err = slash.RegisterCommands()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Success("Slash commands created!")

	stop := make(chan os.Signal, 1)
	//lint:ignore SA1016 i dont know, it just works lol
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	<-stop
	err = discord.Disconnect()
	if err != nil {
		logger.Error("Cannot disconnect... we are disconnecting anyway...", err)
	}
	logger.Success("Exiting...")
}
