package dnd5e

import (
	"fmt"
	"github.com/CoyoTan/JACK-AL/bot/responders/dnd5e/gcore"
	"github.com/CoyoTan/JACK-AL/config"
	"github.com/CoyoTan/JACK-AL/structs"
	"google.golang.org/api/calendar/v3"
	"os"
)

var (
	Jackal *structs.CoreCfg

	DndCore = DndConf{
		Version:       "1.0",
		DndWorkingDir: config.GetConfDir() + "/dnd5e",
	}
)

type DndConf struct {
	Version string `json:"-"`

	DndWorkingDir string `json:"-"`

	GCore *calendar.Service `json:"-"`
}

func (d *DndConf) GetGCore() *calendar.Service {
	return d.GCore
}

func (d *DndConf) GetDndDir() string {
	return d.DndWorkingDir
}

func init() {
	fmt.Println("Starting DND5e Module")
}

func InitDnd(core *structs.CoreCfg) (err error) {
	Jackal = core
	Jackal.Logger.Info.Println("Initializing DND5E Module")
	LoadDndCFG("dndConfig.json", &DndCore)
	gcore.InitGoogleCore(jackal, DndCore)
	//FIXME: Placeholder error.
	return nil
}

func LoadDndCFG(fName string, conf *DndConf) {

	if _, err := os.Stat(fName); os.IsNotExist(err) {
		err := os.MkdirAll(DndCore.DndWorkingDir+"/", 660)

		if err != nil {
			Jackal.Logger.Error.Println("DND5E Mod Non-Fatal Error: Failed to make all directories", err)
			Jackal.Logger.Error.Println("DND5E Mod: Persistence will not be enabled")
		}

		newConf := DndConf{}

		err = config.SaveCfg(DndCore.DndWorkingDir+"/"+fName, newConf)

		if err != nil {
			Jackal.Logger.Error.Println("DND5E Mod Non-Fatal Error: Failed to create Config", err)
			Jackal.Logger.Error.Println("DND5E Mod: Persistence will not be enabled")
		}
	} else {
		err := config.LoadCfg(DndCore.DndWorkingDir+"/"+fName, conf)

		if err != nil {
			Jackal.Logger.Error.Println("There was a critical error loading the DND5E Configuration.", err)
		}
	}
	return
}

func SaveDndCFG(fName string, conf *DndConf) {
	err := config.SaveCfg(DndCore.DndWorkingDir+"/"+fName, &conf)

	if err != nil {
		Jackal.Logger.Error.Println("DND5E Mod. Error saving config file!", err)
	}
}
