package help

import (
	"eln.bot/variables"
	"github.com/bwmarrin/discordgo"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == variables.BotId {
		return
	}

	if m.Content == variables.Config.BotPrefix+"help" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Funcionei")
	}

}
