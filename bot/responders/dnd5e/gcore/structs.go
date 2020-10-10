package gcore

import "google.golang.org/api/calendar/v3"

type dndCore interface {
	SetGCore(core *calendar.Service)
	GetGCore() *calendar.Service
	GetDndDir() string
	GetGuildMap() map[string]string
	GetGuildCalendars() []GuildCalender
}

type GuildCalender interface {
	SetCalenderID(calendarID string)
	SetGuildID(guildID string)
}
