package dnd5e

import (
	"google.golang.org/api/calendar/v3"
)

type e5Conf struct {
	Version string `json:"-"`

	DndWorkingDir string `json:"-"`

	GCore *calendar.Service `json:"-"`

	DndGuildCalendars string `json:"DndCalendars"`
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

func (d *e5Conf) SetDndGuildCal(DndGuildCalendersJSON string) {
	d.DndGuildCalendars = DndGuildCalendersJSON
}

func (d *e5Conf) GetDndGuildCal() string {
	return d.DndGuildCalendars
}
