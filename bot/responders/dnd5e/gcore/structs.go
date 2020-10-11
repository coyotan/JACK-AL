package gcore

import (
	"errors"
	"github.com/CoyoTan/JACK-AL/structs"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type e5GCore struct {
	dCore dndCore `json:"-"`

	DndGuildCalendars []CalendarsByGuild        `json:"DndGuildCalenders"`
	DndCalendarMap    map[string]string         `json:"-"`
	GuildCalendars    map[string]*GuildCalender `json:"-"`
}

func (d *e5GCore) AddGuildCalToMap(guildID string, calendarID string) {

	update := false

	for _, v := range d.DndGuildCalendars {
		if v.GuildID == guildID {
			update = true
			v.CalenderID = calendarID
		}
		return
	}

	if !update {
		d.DndGuildCalendars = append(d.DndGuildCalendars, CalendarsByGuild{
			GuildID:    guildID,
			CalenderID: calendarID,
		})
	}
}

func (d *e5GCore) GenerateCalenderMap() {

	//Wipe it!
	d.DndCalendarMap = nil
	//Regenerate it1
	d.DndCalendarMap = make(map[string]string)

	for _, v := range d.DndGuildCalendars {
		d.DndCalendarMap[v.GuildID] = v.CalenderID
	}
}

//CalendarsByGuild provides an interface which can be marshaled into JSON format and saved into the DND5e config file. It contains two string values which can be used to accurately identify which DND Game is currently being serviced.
type CalendarsByGuild struct {
	CalenderID string `json:"CalendarID"`
	GuildID    string `json:"GuildID"`
}

/* Example:
{
	"title" : "Marches of Tevian",
	"Description" : "Marches of Tevian is a Westmarch style DND game based on the 5th Editition. This calendar is designed to sync up the events of the game with the bot, JACK-AL, for announcement and general gameplay services.",
}*/
//GuildCalendar acts an an overarching parent structure of the Daily and Special events. It ties directly to the JSON data stored in the description of a properly configured DND5e Calendar
type GuildCalender struct {
	GuildID string `json:"DiscordGuildID"`
	Desc    string `json:"Description"`

	DailyEvt DailyCalendarEvent `json:"-"`
	SpecEvts []SpecialEvent     `json:"-"`
}

/*
Each day where the bot is expected to perform daily actions should be marked with a Daily Event. Inside the daily event, there is a list of optional functions which the bot may be instructed to do. This includes
activities such as Restocking shops, providing daily updates (read as "News"), inform players of upcoming events (and request RSVP), or allow the DM to address everyone in the server all at once. Daily Event
configuration is done similar to calendar configuration in the sense that it relies on a marshaled structure in the description of the event. Below is an example of a Daily Event.


Google Calendars Event Title: "Daily Event: Bee and Barley Job Fair Announcement"
{
	"Mentions" : ["everyone","adventurers"],
	"Restock" : ["Hafrins", "Bee and Barley Tavern", "ZeZe"],
	"DailyCondition": ["Random", "King BDay"],
	"UpEvents" : [
		{
			"Title" : "Bee and Barley Tavern Job Fair",
			"Description" : "The Bee and Barley Tavern is pleased to announce that we are hosting an Adventurer's Job Fair! Please feel welcome to stop by. We will have many guests, all of which are looking for the services of very skilled individuals to complete wondrous adventures!",
			"Role" : "event1",
			"Emote": ":event1:",
		}
	],
	"Description" : "Citizens, today is the celebration of the birth of our king! In honor of his birth, the kingdom will pay their taxes to the citizens! Each citizen will receive exactly one week's worth of their taxes back from the kingdom.",
}

Data contained within the "Restock" field will be presented to the players in a message which states which shops have been restocked. This field is not required.
The data contained in the "UpEvent" section will be announced in a separate message following the ping created by the first message. This field is not required.
The description of the overall daily event will be put into the initial message where the users are pinged, along with the Restock notifications. This field is required.
*/
//DailyCalendarEvent is the structure that receives data from the description from events named "Daily Event" in the Google Calendar which represents the DND Session.
type DailyCalendarEvent struct {
	Mentions    []string       `json:"Mentions,omitempty"`
	Restock     []string       `json:"Restock,omitempty"`
	Condition   []string       `json:"DailyCondition,omitempty"`
	UpEvents    []UpEvent      `json:"UpcomingEvents,omitempty"`
	Description string         `json:"Description"`
	GuildCal    *GuildCalender `json:"-"`
}

/*
UpEvent is a structure that contains information on upcoming events. This data is read from the Daily Event that is retrieved from the DND Game Calendar.
All fields in this structure are required, if an UpEvent is scheduled. UpEvents must eventually be cleared by the triggering of the scheduled event.

This information is an announcement to the user, not an actual event.

{
	"Title" : "Bee and Barley Tavern Job Fair",
	"Description" : "The Bee and Barley Tavern is pleased to announce that we are hosting an Adventurer's Job Fair! Please feel welcome to stop by. We will have many guests, all of which are looking for the services of very skilled individuals to complete wondrous adventures!",
	"Role" : "event1",
	"Emote": ":event1:",
}

Title is the name of the upcoming event.
Description describes the upcoming event to the player.
Role is the role which will be provided to players that react to the message with an emote.
Emote is the emote that the players much react to the message with to receive the role.
*/
type UpEvent struct {
	Name  []string            `json:"Title"`
	Desc  []string            `json:"Description"`
	Role  []string            `json:"Role"`
	Emote string              `json:"Reaction"`
	Daily *DailyCalendarEvent `json:"-"`
}

/*Example
{
	"Title" : "Bee and Barley Job Fair",
	"Description" : "The Bee and Barley Tavern Job Fair has started! Visitors within the first 30 minutes are provided with a free round of their choice!",
	"Role" : "event1",
}

Title will be provided as the title of the embed message sent in the notifications channel.
Description will be the content of the message that is sent int he notifications channel.
Role will be the role that is mentioned prior to the embed in the notifications channel to get the awareness of the participants.
*/
/*
SpecialEvent is a structure which holds information relating to the events that will occur in the calendar date of the respective day.
These structures are created the day an event is to be fired and populated by the function GetDailyEventSummary. A watch timer is then created and waits for the proper time to notify users that a special event
has begun.
*/
type SpecialEvent struct {
	Name     []string       `json:"Title"`
	Desc     []string       `json:"Description"`
	Role     string         `json:"Role"`
	GuildCal *GuildCalender `json:"-"`
}

//DiscordRole attempts to resolve the Discord Role ID for the role that is provided to the SpecialEvent.
func (e *SpecialEvent) DiscordRole(jackal *structs.CoreCfg) (role *discordgo.Role, err error) {
	roles, err := jackal.Discord.Session.GuildRoles(e.GuildCal.GuildID)

	if err != nil {
		err = errors.New("Failed to get DiscordGo Roles for Guild " + e.GuildCal.GuildID)
	}

	for _, v := range roles {
		if strings.ToLower(v.Name) == e.Role || e.Role == v.ID {
			role = v
		}
	}

	if role == nil {
		err = errors.New("could not find role by name or ID " + e.Role + " in " + e.GuildCal.GuildID)
	}

	return
}
