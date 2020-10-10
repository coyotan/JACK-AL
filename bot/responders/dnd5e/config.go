package dnd5e

import (
	"fmt"
	"github.com/CoyoTan/JACK-AL/bot/responders/dnd5e/gcore"
	"github.com/CoyoTan/JACK-AL/config"
	"github.com/CoyoTan/JACK-AL/structs"
	"os"
)

var (
	jackal *structs.CoreCfg

	e5Core = e5Conf{
		Version:       "1.0",
		DndWorkingDir: config.GetConfDir() + "/dnd5e",
	}
)

func init() {
	fmt.Println("Starting DND5e Module")
}

func InitDnd(core *structs.CoreCfg) (err error) {
	jackal = core
	jackal.Logger.Info.Println("Initializing DND5E Module")
	LoadDndCFG("dndConfig.json", &e5Core)
	gcore.InitGoogleCore(jackal, &e5Core)
	//FIXME: Placeholder error.
	return nil
}

func LoadDndCFG(fName string, conf *e5Conf) {

	if _, err := os.Stat(fName); os.IsNotExist(err) {
		err := os.MkdirAll(e5Core.DndWorkingDir+"/", 660)

		if err != nil {
			jackal.Logger.Error.Println("DND5E Mod Non-Fatal Error: Failed to make all directories", err)
			jackal.Logger.Error.Println("DND5E Mod: Persistence will not be enabled")
		}

		newConf := e5Conf{}

		err = config.SaveCfg(e5Core.DndWorkingDir+"/"+fName, newConf)

		if err != nil {
			jackal.Logger.Error.Println("DND5E Mod Non-Fatal Error: Failed to create Config", err)
			jackal.Logger.Error.Println("DND5E Mod: Persistence will not be enabled")
		}
	} else {
		err := config.LoadCfg(e5Core.DndWorkingDir+"/"+fName, conf)

		if err != nil {
			jackal.Logger.Error.Println("There was a critical error loading the DND5E Configuration.", err)
		}
	}
	return
}

func SaveDndCFG(fName string, conf *e5Conf) {
	err := config.SaveCfg(e5Core.DndWorkingDir+"/"+fName, &conf)

	if err != nil {
		jackal.Logger.Error.Println("DND5E Mod. Error saving config file!", err)
	}
}

func AddCalBuBuild(guildID string, calenderID string) {
	e5Core.AddGuildCalToMap(guildID, calenderID)
	e5Core.GenerateCalenderMap()
}
