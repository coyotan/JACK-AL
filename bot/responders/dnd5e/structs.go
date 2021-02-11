package dnd5e

import (
	"errors"
	"github.com/bwmarrin/discordgo"
)

//Add custom/predefined Header and Footer.
type BasicNotif struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Picture     string `json:"Picture, omitempty"`
	Color       string `json:"Color, omitempty"`
}

//Validate ensures that the fields are properly populated, to prevent errors from occurring when we attempt to generate an embedded message from this.
func (b *BasicNotif) Validate() (valid bool, reason error) {
	if len(b.Title) == 0 {
		return false, errors.New("title cannot be empty")
	}
	return true, nil
}

///TODO: Test this function.
func (b *BasicNotif) GenerateEmbed() (embed *discordgo.MessageEmbed, err error) {

	if valid, reason := b.Validate(); valid {
		embed.Title = b.Title
		embed.Description = b.Description
		embed.Color, err = toIntColor(hexaNumberToInteger(b.Color))
		embed.Image = &discordgo.MessageEmbedImage{URL: b.Picture}
		embed.Footer = &discordgo.MessageEmbedFooter{
			Text:    "JACK-AL Framework | DND5e Extension.",
			IconURL: "https://cloud.coyotech.dev/index.php/s/7jsto99sqHzpi9P/preview",
		}
		return embed, nil
	} else {
		return nil, reason
	}
}

//Add custom/predefined Header and Footer.
type QuestBoard struct {
	Title       string   `json:"Title"`
	Description string   `json:"Description"`
	Picture     string   `json:"Picture, omitempty"`
	Color       string   `json:"Color, omitempty"`
	QuestID     string   `json:"QuestID"`
	Cost        []string `json:"Cost, omitempty"`
	Reward      []string `json:"Reward, omitempty"`
	Difficulty  string   `json:"Difficulty"`
	Location    string   `json:"Location"`
	Origin      string   `json:"Origin"`
	PartySize   string   `json:"PartySize"`
	React       string   `json:"React, omitempty"`
	QuestRole   string   `json:"QuestRole, omitempty"`
	Mentions    []string `json:"Mentions, omitempty"`
}

//Validate ensures that the fields are properly populated, to prevent errors from occurring when we attempt to generate an embedded message from this.
func (q *QuestBoard) Validate() (valid bool, reason error) {
	if len(q.Title) == 0 || len(q.Description) == 0 || len(q.QuestID) == 0 || len(q.Location) == 0 || len(q.Origin) == 0 || len(q.Difficulty) == 0 || len(q.PartySize) == 0 {
		return false, errors.New("fields incorrect. Title, Description, QuestID, Location, Origin, Difficulty, and PartySize are all required fields")
	} else if (len(q.React) > 0 && len(q.QuestRole) == 0) || (len(q.QuestRole) > 0 && len(q.React) == 0) {
		return false, errors.New("if either React or QuestRole are populated, then both are required fields")
	}
	return true, nil
}

///TODO: Make this one a complicated one. This one should have way more to it.
func (q *QuestBoard) GenerateEmbed() (embed *discordgo.MessageEmbed, err error) {
	if valid, reason := q.Validate(); valid {
		embed.Title = q.Title
		embed.Description = q.Description
		embed.Color, err = toIntColor(hexaNumberToInteger(q.Color))
		embed.Image = &discordgo.MessageEmbedImage{URL: q.Picture}
		embed.Footer = &discordgo.MessageEmbedFooter{
			Text:    "JACK-AL Framework | DND5e Extension.",
			IconURL: "https://cloud.coyotech.dev/index.php/s/7jsto99sqHzpi9P/preview",
		}
		return embed, nil
	} else {
		return nil, reason
	}
}

//Add custom/predefined Header and Footer.
type Daily struct {
	Role []string `json:"Role, omitempty"`
	BasicNotif
}

///TODO: Make this one ping people.
func (d *Daily) GenerateEmbed() (embed *discordgo.MessageEmbed, err error) {

	if valid, reason := d.Validate(); valid {
		embed.Title = d.Title
		embed.Description = d.Description
		embed.Color, err = toIntColor(hexaNumberToInteger(d.Color))
		embed.Image = &discordgo.MessageEmbedImage{URL: d.Picture}
		embed.Footer = &discordgo.MessageEmbedFooter{
			Text:    "JACK-AL Framework | DND5e Extension.",
			IconURL: "https://cloud.coyotech.dev/index.php/s/7jsto99sqHzpi9P/preview",
		}
		return embed, nil
	} else {
		return nil, reason
	}
}

type Taxes struct {
	Announced   bool   `json:"Announced"`
	Title       string `json:"Title, omitempty"`
	Description string `json:"Description, omitempty"`
	Color       string `json:"Color, omitempty"`
	Percentage  int    `json:"Percentage, omitempty"`
	Amount      int    `json:"Amount, omitempty"`
}

//Validate ensures that the fields are properly populated, to prevent errors from occurring when we attempt to generate an embedded message from this.
func (t *Taxes) Validate() (valid bool, reason error) {
	if t.Announced {
		if len(t.Title) == 0 {
			return false, errors.New("announced is true. title cannot be empty")
		}
	} else if t.Percentage == 0 && t.Amount == 0 {
		return false, errors.New("percentage and amount are both null. this is either a mistake or this event is useless")
	}
	return true, nil
}

///TODO: Test this function.
func (t *Taxes) GenerateEmbed() (embed *discordgo.MessageEmbed, err error) {

	if valid, reason := t.Validate(); valid {
		embed.Title = t.Title
		if len(t.Description) > 1 {
			embed.Description = t.Description
		}
		if len(t.Color) > 1 {
			embed.Color, err = toIntColor(hexaNumberToInteger(t.Color))
		} else {
			embed.Color = 0x00c5ff
		}

		embed.Footer = &discordgo.MessageEmbedFooter{
			Text:    "JACK-AL Framework | DND5e Extension.",
			IconURL: "https://cloud.coyotech.dev/index.php/s/7jsto99sqHzpi9P/preview",
		}

		return embed, nil
	} else {
		return nil, reason
	}
}

type IncomePayment struct {
	Announced   bool   `json:"Announced"`
	Title       string `json:"Title, omitempty"`
	Description string `json:"Description, omitempty"`
	Color       string `json:"Color, omitempty"`
	IncomeTax   int    `json:"IncomeTax, omitempty"`
}

//Validate ensures that the fields are properly populated, to prevent errors from occurring when we attempt to generate an embedded message from this.
func (i *IncomePayment) Validate() (valid bool, reason error) {
	if i.Announced {
		if len(i.Title) == 0 {
			return false, errors.New("announced is true. title cannot be empty")
		}
	}
	return true, nil
}

func (i *IncomePayment) GenerateEmbed() (embed *discordgo.MessageEmbed, err error) {
	if valid, reason := i.Validate(); valid {
		embed.Title = i.Title
		if len(i.Description) > 1 {
			embed.Description = i.Description
		}
		if len(i.Color) > 1 {
			embed.Color, err = toIntColor(hexaNumberToInteger(i.Color))
		} else {
			embed.Color = 0x00c5ff
		}

		embed.Footer = &discordgo.MessageEmbedFooter{
			Text:    "JACK-AL Framework | DND5e Extension.",
			IconURL: "https://cloud.coyotech.dev/index.php/s/7jsto99sqHzpi9P/preview",
		}
		return embed, nil
	} else {
		return nil, reason
	}
}
