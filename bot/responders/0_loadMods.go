package responders

import (
	"../../structs"
	"errors"
	"github.com/bwmarrin/discordgo"
)

var (
	jackal              *structs.CoreCfg
	createLocalListener map[string]func(message *discordgo.Message) (err error)
)

func init() {
	createLocalListener = make(map[string]func(message *discordgo.Message) (err error))
}

//InitAll starts the initialization process when called on by another package. In this case, bot, the package just above us.
func InitAll(core *structs.CoreCfg) {
	jackal = core
	jackal.Logger.Info.Println("Beginning module load process.")
	jackal.Discord.CreateListeners = createLocalListener
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
