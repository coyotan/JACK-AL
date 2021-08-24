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

//TODO:
//CreateUserRecord needs testing
func (database *Db) CreateUserRecord(userid string, email string, username string, avatar string, locale string, discriminator string, publicFlags int, isdeveloper bool, isadmin bool, bot bool, verified bool, mfaenabled bool) (err error) {
	if err := database.session.Query(`INSERT INTO jackal.users (userid, email, username, avatar, locale, discriminator, publicflags, isdeveloper, isadmin, verified, mfaenabled, bot) VALUES (? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? )`, userid, email, username, avatar, locale, discriminator, publicFlags, isdeveloper, isadmin, verified, mfaenabled, bot).Exec(); err != nil {
		return err
	}
	return
}

func (database *Db) SelectUserByID(userid string) {
	var id gocql.UUID
	var allRemainingData string

	scanner := database.session.Query(`SELECT * FROM jackal.users WHERE userid = ? `, userid).Iter().Scanner()

	for scanner.Next() {
		var (
			userid        string
			avatar        string
			bot           bool
			discriminator string
			email         string
			isAdmin       bool
			isDeveloper   bool
			locale        string
			publicFlags   int
			username      string
			mfa           bool
			verified      bool
		)
		//Note, apparently the database prefers to return shit in alphabetical order EXCEPT the primary key. We will have to keep this in mind when dealing with queries.
		err := scanner.Scan(&userid, &avatar, &bot, &discriminator, &email, &isAdmin, &isDeveloper, &locale, &mfa, &publicFlags, &username, &verified)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(userid, email, username, avatar, locale, discriminator, publicFlags, isDeveloper, isAdmin, verified, mfa, bot)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(id, allRemainingData)

}

//TODO: Add AlterMessage to take into account for updates and edits to messages.
func (database *Db) AddMessage(message *discordgo.Message) (err error) {

	messageJson, err := json.Marshal(message)

	if err != nil {
		return err
	}

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

	return
}
