package bot

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Get text input from the terminal. Return a string so that the terminal loop can process submitted text.
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
