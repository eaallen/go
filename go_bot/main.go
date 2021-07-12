package main

import (
	"fmt"

	"github.com/aichaos/rivescript-go"
)

func main() {
	// Create a new bot with the default settings.
	bot := rivescript.New(nil)

	// To enable UTF-8 mode, you'd have initialized the bot like:
	// bot = rivescript.New(rivescript.WithUTF8())

	// Load a directory full of RiveScript documents (.rive files)
	err := bot.LoadDirectory("C:/rivescript_templates/chat")
	if err != nil {
		fmt.Printf("Error loading from directory: %s", err)
	}

	// Load an individual file.
	// err = bot.LoadFile("./testsuite.rive")
	// if err != nil {
	//   fmt.Printf("Error loading from file: %s", err)
	// }

	// Sort the replies after loading them!
	bot.SortReplies()

	// Get a reply.
	reply, err := bot.Reply("local-user", "my name is eli you are bot")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("The bot says: %s", reply)
	}
}
