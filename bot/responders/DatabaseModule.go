package responders

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("Initializing Jackal Database module.")

	if err := jackal.InitDB(); err != nil {
		jackal.Logger.Error.Println("There was a critical error opening the Jackal Database.", err.Error())
		os.Exit(20)
	}
}
