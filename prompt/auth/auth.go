package authPrompt

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"github.com/lepra-tsr/gdbt/util"
	"golang.org/x/crypto/ssh/terminal"
)

type Auth struct {
	Email    string
	Password string
}

func (u *Auth) AskEmail() error {
	fmt.Println("e-mail: ")
	buf := bufio.NewReader(os.Stdin)
	bytes, err := buf.ReadBytes('\n')
	if err != nil {
		return err
	}
	u.Email = util.StripNewLine(string(bytes))

	return nil
}

func (u *Auth) AskPassword() error {
	fmt.Println("password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		// terminal.ReadPassword は、MINGW系(windows)のターミナルソフトで使用できない: ターミナルソフト側の問題らしい)
		fmt.Println("* sorry, your shell CANNOT hide password :<")
		buf := bufio.NewReader(os.Stdin)
		bytePassword, err := buf.ReadBytes('\n')
		if err != nil {
			return err
		}
		u.Password = util.StripNewLine(string(bytePassword))
		return nil
	}
	u.Password = util.StripNewLine(string(bytePassword))
	return nil
}
