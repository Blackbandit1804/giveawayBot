package main

import (
	//"fmt"
	"strings"
	"errors"
	"io/ioutil"
	"encoding/json"
	"github.com/bwmarrin/discordgo"
)

//////////////// PUBLIC STUFF ////////////////

// SendEmbed sends a prepared embed message with content passed by arguments in passed channel
func SendEmbed(session *discordgo.Session, channelID string, cont string) (*discordgo.Message, error) {
	return session.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
		Description: cont,
		Color: COLOR_MAIN,
	})
}

// SendEmbed sends a prepared error type embed message with content passed by arguments in passed channel
func SendEmbedError(session *discordgo.Session, channelID string, cont string) (*discordgo.Message, error) {
	return session.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
		Description: cont,
		Color: COLOR_ERROR,
	})
}

// CheckAutorized returns if the passed user has one of
// the authorized roles
func CheckAutorized(config *Config, guildid string, member *discordgo.Member) bool {
	authRoles := map[string][]string{}
	bData, err := ioutil.ReadFile("./.authroles.json")
	if err == nil {
		err = json.Unmarshal(bData, &authRoles)
	}

	for _, r := range member.Roles {
		for _, a := range authRoles[guildid] {
			if r == a {
				return true
			}
		}
	}
	return false
}

// CheckAdmin returns if the passed user has the same ID as the entered
// admin ID in the config
func CheckAdmin(config *Config, user *discordgo.User) bool {
	return config.Data.Admin == user.ID
}

// FetchChannel returns a text channel fetched by pased ID, mention, name or name-part
func FetchChannel(g *discordgo.Guild, resolvable string) (*discordgo.Channel, error) {
	var channels []*discordgo.Channel
	for _, c := range g.Channels {
		if c.Type == discordgo.ChannelTypeGuildText {
			channels = append(channels, c)
		}
	} 
	for _, c := range channels {
		if c.ID == strings.Replace(resolvable, "<#", ">", -1) {
			return c, nil
		}
	}
	for _, c := range channels {
		if strings.ToLower(c.Name) == strings.ToLower(resolvable) {
			return c, nil
		}
	}
	for _, c := range channels {
		if strings.ToLower(c.Name) == strings.ToLower(resolvable) {
			return c, nil
		}
	}
	for _, c := range channels {
		if strings.HasPrefix(strings.ToLower(c.Name), strings.ToLower(resolvable)) {
			return c, nil
		}
	}
	return nil, errors.New("channel not found")
}

func FetchRole(g *discordgo.Guild, resolvable string) (*discordgo.Role, error) {
	roles := g.Roles 
	for _, c := range roles {
		if c.ID == strings.Replace(resolvable, "<@&", ">", -1) {
			return c, nil
		}
	}
	for _, c := range roles {
		if strings.ToLower(c.Name) == strings.ToLower(resolvable) {
			return c, nil
		}
	}
	for _, c := range roles {
		if strings.ToLower(c.Name) == strings.ToLower(resolvable) {
			return c, nil
		}
	}
	for _, c := range roles {
		if strings.HasPrefix(strings.ToLower(c.Name), strings.ToLower(resolvable)) {
			return c, nil
		}
	}
	return nil, errors.New("role not found")
}