package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/coyotan/JACK-AL/bot/responders"
	"github.com/coyotan/JACK-AL/botutils"
	"github.com/coyotan/JACK-AL/structs"
	"github.com/txgruppi/parseargs-go"
	"os"
)

var (
	jackal *structs.CoreCfg
)

//Init accepts a pointer to the core of Jackal. This core will be used to establish a connection with Discord and act as a registration point for command handlers, as well as a service provider for other linked services.
func Init(core *structs.CoreCfg) {

	jackal = core

	jackal.Logger.Info.Println("Main process has handed off &Core to package: bot")

	dgOpen()

	for {

		resp, _ := parseargs.Parse(GetInput())

		if len(resp) > 0 {
			switch resp[0] {
			case "ping":
				jackal.Logger.Console.Println("Pong")
			case "leave":

				err := botutils.SaveCfg(botutils.ConfigDir+"config.json", jackal)

				if err != nil {
					jackal.Logger.Console.Println("Failed to save to config.json. Current running configuration will be lost. Please check the log file for more information.")
					jackal.Logger.Error.Println(err)
					jackal.Logger.Info.Println("JACK-AL closing with errors. Exit Code 199")
					jackal.Discord.Session.Close()
					os.Exit(199)
				} else {
					jackal.Logger.Info.Println("Peacefully closing JACK-AL. Exit Code 100")
					jackal.Discord.Session.Close()
					os.Exit(100)
				}
			}
		}
	}
}

func dgOpen() {

	if !(len(jackal.Discord.Token) > 0) {
		jackal.Logger.Error.Println("The configuration provided does not contain an API Token. Please provide a token to the Jackal Configuration file.", len(jackal.Discord.Token))
	}

	dg, err := discordgo.New("Bot " + jackal.Discord.Token)

	//Enable and configure Stateful Discord!
	dg.StateEnabled = true
	dg.State.TrackRoles = true
	dg.State.TrackMembers = true
	dg.State.TrackPresences = true

	if jackal.Discord.MaxMessageCount < 1 {
		dg.State.MaxMessageCount = 50
	} else {
		dg.State.MaxMessageCount = jackal.Discord.MaxMessageCount
	}

	if err != nil {
		jackal.Logger.Error.Println("There was an error when attempting to begin a session with Discord.")
	}

	responders.InitAll(jackal)

	dg.AddHandler(ready)
	dg.AddHandler(createDispatch)
	dg.AddHandler(nonprefixDispatch)
	dg.AddHandler(editDispatch)
	//dg.AddHandler(deleteDispatch)
	dg.AddHandler(addReactionDispatch)
	dg.AddHandler(rmReactionDispatch)

	initModDispatch(jackal)

	err = dg.Open()

	if err != nil {
		jackal.Logger.Console.Println(err)
		jackal.Logger.Error.Println(err)
		os.Exit(101)
	}

	if err := jackal.InitCassandraDB(); err != nil {

		if jackal.IsDockerContainer() {
			fmt.Println("Is a docker container!")
			//Announce that we will configure the database for them!
			jackal.Logger.Console.Println("JACKAL is currently configuring the Cassandra database. Please standby.")

			fmt.Println(jackal.Database.CreateUserTable())

		} else {
			//Tell them that since they didn't use the docker container, they need to do it themselves!
			jackal.Logger.Console.Println("JACK-AL has detected that it is not in a docker container, and that the database has not yet been configured. JACK-AL cannot automatically configure databases at this time. Please follow the guidance on the github page to manually configure the Cassandra database.")
			//TODO: Make this not matter in the future. JACKAL should be able to configure the database, so long as it knows where to find it. This snip of code probably won't even make it to the next release, but it is here for now.
			jackal.Logger.Error.Println("There was a critical error opening the Jackal Database.", err.Error())
			os.Exit(20)
		}
	}

	jackal.Logger.Info.Println("JackalDB Initialization Completed!")

}
