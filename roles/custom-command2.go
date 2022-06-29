package roles

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/shinPallini/discordgox"
)

func init() {
	embed := discordgox.NewMessageEmbed(
		discordgox.SetTitle("Embed_New"),
		discordgox.SetDescription("New Embed_New"),
		discordgox.SetColor(0x111111),
		discordgox.SetEmbedField(
			discordgox.NewList(
				discordgox.NewMessageEmbedField(
					discordgox.SetEmbedFieldName("field11"),
					discordgox.SetEmbedFieldValue("value11"),
				),
			),
		),
	)
	response := discordgox.NewInteractionResponse(
		discordgox.SetType(discordgo.InteractionResponseChannelMessageWithSource),
		discordgox.SetData(discordgox.NewInteractionResponseData(
			discordgox.SetContent("custom interaction2"),
			discordgox.SetEmbed(
				discordgox.NewList(
					discordgox.NewMessageEmbed(
						discordgox.SetTitle("Embed"),
						discordgox.SetDescription("New Embed"),
						discordgox.SetColor(0x111111),
						discordgox.SetEmbedField(
							discordgox.NewList(
								discordgox.NewMessageEmbedField(
									discordgox.SetEmbedFieldName("field1"),
									discordgox.SetEmbedFieldValue("value1"),
								),
							),
						),
					),
					embed,
				),
			),
			discordgox.SetComponent(
				discordgox.NewList[discordgo.MessageComponent](
					discordgox.NewActionsRow(
						discordgox.SetLinkButton(
							"Custom link",
							"https://go.dev/doc/tutorial/call-module-code",
						),
					),
				),
			),
		)),
	)

	discordgox.AddCommand(
		&discordgo.ApplicationCommand{
			Name:        "custom-command2",
			Description: "Custom command2",
		},
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, response)
		},
	)
	log.Println(response.Data)
	log.Println(discordgox.Commands)
}
