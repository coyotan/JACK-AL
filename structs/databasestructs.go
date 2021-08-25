package structs

import "github.com/bwmarrin/discordgo"

type DBUser struct {
	JDB  DBData
	User discordgo.User
}

type DBData struct {
	isAdmin     bool
	isDeveloper bool
}

func (dbd *DBData) GetAdmin() (isAdmin bool) {
	return dbd.isAdmin
}

func (dbd *DBData) GetDeveloper() (isDeveloper bool) {
	return dbd.isDeveloper
}
