package gcore

import "google.golang.org/api/calendar/v3"

type dndCore interface {
	getGCore() *calendar.Service
	getDndDir() string
}
