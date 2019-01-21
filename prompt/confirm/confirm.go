package confirmPrompt

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/lepra-tsr/gdbt/util"
	"os"
	"strings"
)

type Confirm struct {
}

func (u *Confirm) AskIn(commandListString string) (string, error) {
	command := ""
	commandList := strings.Split(commandListString, ",")
	for i := 0; true; i++ {
		buf := bufio.NewReader(os.Stdin)
		bytes, err := buf.ReadBytes('\n')
		if err != nil {
			return "", err
		}
		command = util.StripNewLine(string(bytes))
		found := false
		for j := 0; j < len(commandList); j++ {
			found = commandList[j] == command
			if found {
				break
			}
		}
		if found {
			break
		} else {
			fmt.Println("\"" + command + "\" does not match " + commandListString)
		}

		if i > 5 {
			return "", errors.New("exceeded confirm limit")
		}
	}

	return command, nil
}
