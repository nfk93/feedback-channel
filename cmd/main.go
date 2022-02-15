package main

import (
	"feedback-channel/command"
	"feedback-channel/interactions"
	"flag"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/go-resty/resty/v2"
)

const hexEncodedDiscordPubkey = "0ece8e19901d676332e1a109f75efc878f56b6d3f01345d575ab91e510ec442d"

// Variables used for command line parameters
var (
	Token           string
	RegisterCommand bool
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.BoolVar(&RegisterCommand, "register", false, "register the slash command")
	flag.Parse()
}

func registerSlashCommand(cmd command.Command) {
	commandClient := resty.New()
	resp, err := commandClient.R().
		SetHeader("Authorization", "Bot "+Token).
		SetHeader("Content-Type", "application/json").
		SetBody(cmd).
		Post("https://discord.com/api/v8/applications/939540521897058304/commands")
	if err != nil {
		panic("shit")
	}

	fmt.Printf("from register command: %s\n", resp.String())
}

type ImplResonse struct {
	Code int
	Body interface{}
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	if RegisterCommand {
		cmd := command.Command{
			Name:        "blep",
			Type:        1,
			Description: "bla",
			Options: []command.CommandOption{
				{
					Name:        "animal",
					Description: "type of animal",
					Type:        3,
					Required:    true,
					Choices: []command.CommandOptionChoice{
						{
							Name:  "dog",
							Value: "animal_dog",
						},
						{
							Name:  "cat",
							Value: "animal_cat",
						},
					},
				},
				{
					Name:        "only_smol",
					Description: "whether o show only smol animals",
					Type:        5,
					Required:    false,
				},
			},
		}
		registerSlashCommand(cmd)
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Just like the ping pong example, we only care about receiving message
	// events in this example.
	dg.Identify.Intents = discordgo.IntentsDirectMessages + discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	interactionService := interactions.New(hexEncodedDiscordPubkey)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	interactionService.Start()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
//
// It is called whenever a message is created but only when it's sent through a
// server as we did not request IntentsDirectMessages.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Type == 2 {
		fmt.Println("RECEIVED TYPE 2")
		fmt.Printf("m: %v\n", m)
		return
	}

	// In this example, we only care about messages that are "ping".
	if m.Content != "ping" {
		return
	}

	// We create the private channel with the user who sent the message.
	channel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		// If an error occurred, we failed to create the channel.
		//
		// Some common causes are:
		// 1. We don't share a server with the user (not possible here).
		// 2. We opened enough DM channels quickly enough for Discord to
		//    label us as abusing the endpoint, blocking us from opening
		//    new ones.
		fmt.Println("error creating channel:", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Something went wrong while sending the DM!",
		)
		return
	}
	// Then we send the message through the channel we created.
	msg := fmt.Sprintf(
		`Received ping from %s
			server: %s
			type: %d`,
		m.Author.Username,
		m.GuildID,
		m.Type)

	_, err = s.ChannelMessageSend(channel.ID, msg)
	if err != nil {
		// If an error occurred, we failed to send the message.
		//
		// It may occur either when we do not share a server with the
		// user (highly unlikely as we just received a message) or
		// the user disabled DM in their settings (more likely).
		fmt.Println("error sending DM message:", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Failed to send you a DM. "+
				"Did you disable DM in your privacy settings?",
		)
	}
}
