package bot

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

///TODO: Add support for separate users to sign in and use the terminal. Do it securely (maybe with Discord OAUTH support?)
//TODO: Provide a console list of servers to select from. Allow for bot-feature management based on selected servers.
//GetInput from the terminal. Return a string so that the terminal loop can process submitted text.
func GetInput() string {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print(">")

		response, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err.Error())
		}

		response = strings.TrimSpace(response)
		return strings.ToLower(response)
	}
}
