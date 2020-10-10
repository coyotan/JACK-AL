package dnd5e

import (
	"github.com/CoyoTan/JACK-AL/config"
	"github.com/CoyoTan/JACK-AL/structs"
	"golang.org/x/oauth2"
	"os"
)

var (
	DndWorkingDir = config.GetConfDir() + "/dnd5e"

	Jackal *structs.CoreCfg

	DndCore = DndConf{
		Version: "1.0",
	}
)

type DndConf struct {
	Version string `json:"-"`

	GCore *oauth2.Config `json:"-"`
}

func init() {
	Jackal.Logger.Info.Println("Initializing DND5E Module")
	LoadDndCFG("dndConfig.json", &DndCore)

}

func InitDnd(core *structs.CoreCfg) {
	Jackal = core
}

func LoadDndCFG(fName string, conf *DndConf) {

	if config.IsFirstRun() {
		err := os.MkdirAll(DndWorkingDir, 660)

		if err != nil {
			Jackal.Logger.Error.Println("DND5E Mod Non-Fatal Error: Failed to make all directories", err)
			Jackal.Logger.Error.Println("DND5E Mod: Persistence will not be enabled")
		}

		newConf := DndConf{}

		err = config.SaveCfg(DndWorkingDir+fName, newConf)

		if err != nil {
			Jackal.Logger.Error.Println("DND5E Mod Non-Fatal Error: Failed to create Config", err)
			Jackal.Logger.Error.Println("DND5E Mod: Persistence will not be enabled")
		}
	} else {
		err := config.LoadCfg(DndWorkingDir+fName, &DndCore)

		if err != nil {
			Jackal.Logger.Error.Println("There was a critical error loading the DND5E Configuration.", err)
		}
	}
	return
}

func Dnd5eSaveCFG(fName string, conf *DndConf) {
	err := config.SaveCfg(DndWorkingDir+fName, &conf)

	if err != nil {
		Jackal.Logger.Error.Println("DND5E Mod. Error saving config file!", err)
	}
}
