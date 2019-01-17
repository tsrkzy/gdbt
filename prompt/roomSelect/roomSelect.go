package authPrompt

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"github.com/lepra-tsr/gdbt/util"
	"golang.org/x/crypto/ssh/terminal"
)

type RoomSelect struct {
}

func (u *RoomSelect) ChooseRoom() error {
	fmt.Println("Choose Room.")
	
	buf := bufio.NewReader(os.Stdin)
	bytes, err := buf.ReadBytes('\n')
	if err != nil {
		return err
	}
	u.Email = util.StripNewLine(string(bytes))

	return nil
}
