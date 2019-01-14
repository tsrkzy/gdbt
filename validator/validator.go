package validator

import (
	"errors"
	"github.com/lepra-tsr/gdbt/util"
)

func ValidateCommand(commandStr string) error {
	commandList := []string{"init", "channel", "list"}
	i := util.StrIndexOf(commandStr, commandList)
	if i == -1 {
		return errors.New("validation error: invalid command: \"" + commandStr + "\"")
	}

	return nil
}