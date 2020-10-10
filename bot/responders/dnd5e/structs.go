package dnd5e

import (
	"errors"
	"github.com/CoyoTan/JACK-AL/structs"
	"github.com/bwmarrin/discordgo"
	"google.golang.org/api/calendar/v3"
	"strings"
)

type e5Conf struct {
	Version string `json:"-"`

	DndWorkingDir string `json:"-"`

	GCore *calendar.Service `json:"-"`

	DndGuildCalendars []CalendarsByGuild `json:"DndCalendars"`
	DndCalendarMap    map[string]string  `json:"-"`
}

func (d *e5Conf) SetGCore(core *calendar.Service) {
	d.GCore = core
}

func (d *e5Conf) GetGCore() *calendar.Service {
	return d.GCore
}

func (d *e5Conf) GetDndDir() string {
	return d.DndWorkingDir
}

func (d *e5Conf) GetGuildMap() map[string]string {
	return d.DndCalendarMap
}

func (d *e5Conf) AddGuildCalToMap(guildID string, calendarID string) {

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

func (d *e5Conf) GenerateCalenderMap() {

	//Wipe it!
	d.DndCalendarMap = nil
	//Regenerate it1
	d.DndCalendarMap = make(map[string]string)

	for _, v := range d.DndGuildCalendars {
		d.DndCalendarMap[v.GuildID] = v.CalenderID
	}
}

type GuildCalender struct {
	GuildID string `json:"DiscordGuildID"`
	Desc    string `json:"Description"`

	DailyEvt DailyCalendarEvent `json:"-"`
	SpecEvts []SpecialEvent     `json:"-"`
}

type DailyCalendarEvent struct {
	Mentions    []string       `json:"Mentions,omitempty"`
	Restock     []string       `json:"Restock,omitempty"`
	Condition   []string       `json:"DailyCondition,omitempty"`
	UpEvents    []UpEvent      `json:"UpcomingEvents,omitempty"`
	Description string         `json:"Description"`
	GuildCal    *GuildCalender `json:"-"`
}

type UpEvent struct {
	Name  []string            `json:"Title"`
	Desc  []string            `json:"Description"`
	Role  []string            `json:"Role"`
	Emote string              `json:"Reaction"`
	Daily *DailyCalendarEvent `json:"-"`
}

type SpecialEvent struct {
	Name     []string       `json:"Title"`
	Desc     []string       `json:"Description"`
	Role     string         `json:"Role"`
	GuildCal *GuildCalender `json:"-"`
}

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

type CalendarsByGuild struct {
	CalenderID string `json:"CalendarID"`
	GuildID    string `json:"GuildID"`
}
