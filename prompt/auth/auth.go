package authPrompt

import (
	"bufio"
	"fmt"
	"github.com/lepra-tsr/gdbt/util"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"syscall"
)

func AskEmail() (string, error) {
	fmt.Println("e-mail: ")
	buf := bufio.NewReader(os.Stdin)
	if email, err := buf.ReadBytes('\n'); err != nil {
		return "", err
	} else {
		return util.StripNewLine(string(email)), nil
	}
}

func AskPassword() (string, error) {
	fmt.Println("password: ")
	if bytePassword, err := terminal.ReadPassword(int(syscall.Stdin)); err != nil {
		// terminal.ReadPassword は、MINGW系(windows)のターミナルソフトで使用できない: ターミナルソフト側の問題らしい)
		fmt.Println("* sorry, your shell CANNOT hide password :<")
		buf := bufio.NewReader(os.Stdin)
		if bytePassword, err := buf.ReadBytes('\n'); err != nil {
			return "", err
		} else {
			return util.StripNewLine(string(bytePassword)), nil
		}
	} else {
		return util.StripNewLine(string(bytePassword)), nil
	}
}
