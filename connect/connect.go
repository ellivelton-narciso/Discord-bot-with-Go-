package connect

import (
	"eln.bot/commands/NH"
	"eln.bot/commands/help"
	"eln.bot/variables"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func Start() {
	var goBot *discordgo.Session

	goBot, err := discordgo.New("Bot " + variables.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	variables.BotId = u.ID

	goBot.AddHandler(NH.NH)
	goBot.AddHandler(help.Help)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot está funcionando !")
}
