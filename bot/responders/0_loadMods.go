package responders

//DOCUMENTATION:
/*
NO FILES INTENDING TO USE THE DISCORD INTERFACE MAY BE ALPHABETICALLY SUPERIOR TO 0_loadMods.go. This will cause load order issues.
Public Variables accessible from this location:
jackal - Core configuration.
jackal.Discord - Contains Discord configuration
*/

import (
	"errors"
	"github.com/CoyoTan/JACK-AL/structs"
	"github.com/bwmarrin/discordgo"
)

var (
	jackal              *structs.CoreCfg
	initLocalListener   map[string]func(core *structs.CoreCfg) (err error)
	editLocalListener   map[string]func(edited *discordgo.MessageUpdate) (err error)
	createLocalListener map[string]func(message *discordgo.Message) (err error)
	deleteLocalListener map[string]func(deleted *discordgo.MessageDelete) (err error)
	reactLocalListener  map[string]func(deleted *discordgo.MessageReaction) (err error)
)

func init() {
	initLocalListener = make(map[string]func(message *structs.CoreCfg) (err error))
	editLocalListener = make(map[string]func(message *discordgo.MessageUpdate) (err error))
	createLocalListener = make(map[string]func(message *discordgo.Message) (err error))
	deleteLocalListener = make(map[string]func(message *discordgo.MessageDelete) (err error))
	reactLocalListener = make(map[string]func(message *discordgo.MessageReaction) (err error))
}

//InitAll starts the initialization process when called on by another package. In this case, bot, the package just above us.
func InitAll(core *structs.CoreCfg) {
	jackal = core
	jackal.Logger.Info.Println("Beginning module load process.")
	jackal.Discord.EditListeners = editLocalListener
	jackal.Discord.InitModListeners = initLocalListener
	jackal.Discord.CreateListeners = createLocalListener
	jackal.Discord.DeleteListeners = deleteLocalListener
	jackal.Discord.ReactListeners = reactLocalListener
}

func addCreateListener(name string, responder func(message *discordgo.Message) (err error)) (err error) {

	if createLocalListener == nil {
		createLocalListener = make(map[string]func(message *discordgo.Message) (err error))
	}

	//This might be able to be revised in the future. It may not be important for users to be able to name their map keys themselves, so we might be able to use a random ID generator for this in the future.
	if _, ok := createLocalListener[name]; ok {
		err = errors.New("CreateListeners already contains a function with this name. Please pick a different name")

		return
	}

	createLocalListener[name] = responder

	return
}

func addInitModListener(name string, responder func(message *structs.CoreCfg) (err error)) (err error) {

	if initLocalListener == nil {
		initLocalListener = make(map[string]func(message *structs.CoreCfg) (err error))
	}

	//This might be able to be revised in the future. It may not be important for users to be able to name their map keys themselves, so we might be able to use a random ID generator for this in the future.
	if _, ok := createLocalListener[name]; ok {
		err = errors.New("InitListeners already contains a function with this name. Please pick a different name")

		return
	}

	initLocalListener[name] = responder

	return
}

func addEditListener(name string, responder func(message *discordgo.MessageUpdate) (err error)) (err error) {

	if editLocalListener == nil {
		editLocalListener = make(map[string]func(message *discordgo.MessageUpdate) (err error))
	}

	//This might be able to be revised in the future. It may not be important for users to be able to name their map keys themselves, so we might be able to use a random ID generator for this in the future.
	if _, ok := createLocalListener[name]; ok {
		err = errors.New("EditListeners already contains a function with this name. Please pick a different name")

		return
	}

	editLocalListener[name] = responder

	return
}

func addDeleteListener(name string, responder func(message *discordgo.MessageDelete) (err error)) (err error) {

	if deleteLocalListener == nil {
		deleteLocalListener = make(map[string]func(message *discordgo.MessageDelete) (err error))
	}

	//This might be able to be revised in the future. It may not be important for users to be able to name their map keys themselves, so we might be able to use a random ID generator for this in the future.
	if _, ok := createLocalListener[name]; ok {
		err = errors.New("DeleteListeners already contains a function with this name. Please pick a different name")

		return
	}

	deleteLocalListener[name] = responder

	return
}

func addReactListener(name string, responder func(message *discordgo.MessageReaction) (err error)) (err error) {

	if createLocalListener == nil {
		reactLocalListener = make(map[string]func(message *discordgo.MessageReaction) (err error))
	}

	//This might be able to be revised in the future. It may not be important for users to be able to name their map keys themselves, so we might be able to use a random ID generator for this in the future.
	if _, ok := createLocalListener[name]; ok {
		err = errors.New("CreateListeners already contains a function with this name. Please pick a different name")

		return
	}

	reactLocalListener[name] = responder

	return
}
