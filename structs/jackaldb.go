package structs

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gocql/gocql"
)

type Db struct {
	IP      string         `json:"cassandraIP"`
	session *gocql.Session `json:"-"`
}

func (database *Db) CreateKeyspace() (err error) {
	if err := database.session.Query(`CREATE KEYSPACE IF NOT EXISTS jackal WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1'};`).Exec(); err != nil {
		return err
	}
	return
}

func (database *Db) CreateUserTable() (err error) {
	if err = database.session.Query(`CREATE TABLE IF NOT EXISTS jackal.users (UserID text PRIMARY KEY, Email text, Username text, Avatar text, Locale text, Discriminator text, PublicFlags int, IsDeveloper boolean, IsAdmin boolean, Verified boolean, MFAEnabled boolean, Bot boolean) WITH compression = {'class': 'LZ4Compressor', 'chunk_length_in_kb': 64, 'crc_check_chance': 0.5};`).Exec(); err != nil {
		return err
	}
	return
}

func (database *Db) CreateMessagesTable() (err error) {

	if err = database.session.Query(`CREATE TABLE IF NOT EXISTS jackal.messages (messageid text, channelid text, guildid text, authorid text, content text, messagetype int, json text, messagesent timestamp, messageupdated timestamp, PRIMARY KEY (messageid, guildid)) WITH compression = {'class': 'LZ4Compressor', 'chunk_length_in_kb': 64, 'crc_check_chance': 0.5};`).Exec(); err != nil {
		return err
	}

	return
}

//TODO: Conduct following
//CreateUserRecord needs testing
func (database *Db) CreateUserRecord(userid string, email string, username string, avatar string, locale string, discriminator string, publicFlags int, isdeveloper bool, isadmin bool, bot bool, verified bool, mfaenabled bool) (err error) {
	if err := database.session.Query(`INSERT INTO jackal.users (userid, email, username, avatar, locale, discriminator, publicflags, isdeveloper, isadmin, verified, mfaenabled, bot) VALUES (? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? )`, userid, email, username, avatar, locale, discriminator, publicFlags, isdeveloper, isadmin, verified, mfaenabled, bot).Exec(); err != nil {
		return err
	}
	return
}

func (database *Db) SelectUserByID(userid string) (result DBUser, err error) {
	var (
		qUserID       = ""
		avatar        = ""
		bot           = false
		discriminator = ""
		email         = ""
		isAdmin       = false
		isDeveloper   = false
		locale        = ""
		publicFlags   = 0
		username      = ""
		mfa           = false
		verified      = false
	)

	//Note, apparently the database prefers to return shit in alphabetical order EXCEPT the primary key. We will have to keep this in mind when dealing with queries.
	err = database.session.Query(`SELECT * FROM jackal.users WHERE userid = ? `, userid).Scan(&qUserID, &avatar, &bot, &discriminator, &email, &isAdmin, &isDeveloper, &locale, &mfa, &publicFlags, &username, &verified)

	if err != nil {
		fmt.Println("Erroring here")
		return DBUser{}, err
	}

	fmt.Println(qUserID, email, username, avatar, locale, discriminator, publicFlags, isDeveloper, isAdmin, verified, mfa, bot)
	fmt.Println("Passing debug printing.")

	//	if err := scanner.Err(); err != nil {
	//		return DBUser{}, err
	//	}

	result = DBUser{
		JDB: DBData{
			isAdmin:     isAdmin,
			isDeveloper: isDeveloper,
		},
		User: discordgo.User{
			ID:            qUserID,
			Email:         email,
			Username:      username,
			Avatar:        avatar,
			Locale:        locale,
			Discriminator: discriminator,
			Verified:      verified,
			MFAEnabled:    mfa,
			Bot:           bot,
		},
	}

	return
}

func (database *Db) AddMessage(message *discordgo.Message) (err error) {

	var (
		channelid = ""
		guildid   = ""
		authorid  = ""
		content   = ""
	)

	messageJson, err := json.Marshal(message)

	if err != nil {
		return err
	}

	err = database.session.Query(`SELECT channelid, guildid, authorid, content FROM jackal.messages WHERE messageid = ?`, message.ID).Scan(&channelid, &guildid, &authorid, &content)

	if err != nil {
		if !(err.Error() == "not found") {
			return err
		} else {
			//If the error is that it is not found, reset the error! This is expected.
			err = nil
		}
	}

	if !(channelid == message.ChannelID && message.GuildID == message.GuildID && message.Author.ID == authorid && message.Content == content) {

		timestamp1, err := message.Timestamp.Parse()
		if err != nil {
			return err
		}

		err = database.session.Query(`INSERT INTO jackal.messages (messageid, channelid, guildid, authorid, content, messagetype, json, messagesent) VALUES( ?, ?, ?, ?, ?, ?, ?, ?)`,
			message.ID,
			message.ChannelID,
			message.GuildID,
			message.Author.ID,
			message.Content,
			message.Type,
			messageJson,
			timestamp1).Exec()

		if err != nil {
			return err
		}
	}

	return err
}

func (database *Db) AddUserFromMessage(message *discordgo.Message) (err error) {

	if err != nil {
		return err
	}

	if result, err := database.SelectUserByID(message.Author.ID); err != nil {
		//TODO: Remove assumption that this means that the user does not exist. This is a BAD fucking idea. Change to database error type checking.

		err = database.session.Query(`INSERT INTO jackal.users (userid, email, username, avatar, locale, discriminator, publicflags, isdeveloper, isadmin, verified, mfaenabled, bot ) VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			message.Author.ID,
			message.Author.Email,
			message.Author.Username,
			message.Author.Avatar,
			message.Author.Locale,
			message.Author.Discriminator,
			message.Author.PublicFlags,
			false,
			false,
			message.Author.Verified,
			message.Author.MFAEnabled,
			message.Author.Bot).Exec()

		if err != nil {
			return err
		}

	} else if result.User.ID == message.Author.ID {
		if !(message.Author.Username == result.User.Username && message.Author.Discriminator == result.User.Discriminator && message.Author.Avatar == result.User.Avatar) {
			err = database.session.Query(`UPDATE jackal.users SET email = ? , username = ? , avatar = ? , locale = ? , discriminator = ? , publicflags = ? , isdeveloper = ? , isadmin = ? , verified = ? , mfaenabled = ? , bot = ? WHERE userid = ? `,
				message.Author.Email,
				message.Author.Username,
				message.Author.Avatar,
				message.Author.Locale,
				message.Author.Discriminator,
				message.Author.PublicFlags,
				result.JDB.GetDeveloper(),
				result.JDB.GetAdmin(),
				message.Author.Verified,
				message.Author.MFAEnabled,
				message.Author.Bot,
				message.Author.ID).Exec()

			if err != nil {
				return err
			}
		}
	}

	return err
}
