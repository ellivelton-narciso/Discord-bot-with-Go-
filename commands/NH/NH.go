package NH

import (
	"eln.bot/variables"
	"github.com/bwmarrin/discordgo"
)

func NH(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == variables.BotId {
		return
	}

	if m.Content == variables.Config.BotPrefix+"NH" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Bora dominar Los Flores!")
	}

}
