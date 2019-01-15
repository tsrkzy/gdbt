package setup

// init は予約語なので避けた

import (
	"bufio"
	"bytes"
	"fmt"

	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"syscall"

	"github.com/lepra-tsr/gdbt/config"
	"github.com/lepra-tsr/gdbt/util"
	"golang.org/x/crypto/ssh/terminal"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func Handler() error {
	fmt.Println("init handler.")
	// mkdir -p ~/.gdbt
	// fopen("w", ~/.gdbt/config)
	// fopen("w", ~/.gdbt/draft)
	if err := config.CheckConfigFileState(); err != nil {
		return err
	}

	// wait for input email and password
	// fetch accesstoken
	token, err := authPrompt()
	if err != nil {
		return err
	}
	fmt.Println("get access token done.")

	if err := saveToken(token); err != nil {
		return err
	}
	fmt.Println(token + " > ~/.gdbt/config.json")

	// fetch userInfo
	// fetch organizationInfo
	// fetch roomInfo

	// write userName
	// write accessToken
	// write user
	// write organizations
	// write rooms

	// wait for input room selection
	return nil
}

func authPrompt() (string, error) {
	var (
		email    string
		password string
	)
	if _email, err := askEmail(); err != nil {
		return "", err
	} else {
		email = _email
	}

	if _password, err := askPassword(); err != nil {
		return "", err
	} else {
		password = _password
	}

	if token, err := fetchToken(email, password); err != nil {
		return "", err
	} else {
		if err := config.WriteCredential(email, token); err != nil {
			return "", err
		}

		return token, nil
	}
}

func askEmail() (string, error) {
	fmt.Println("e-mail: ")
	buf := bufio.NewReader(os.Stdin)
	if email, err := buf.ReadBytes('\n'); err != nil {
		return "", err
	} else {
		return util.StripNewLine(string(email)), nil
	}
}

func askPassword() (string, error) {
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

func fetchToken(email string, password string) (string, error) {
	url := "https://idobata.io/oauth/token"
	payload := fmt.Sprintf(`{"grant_type":"password","username":"%v","password":"%v"}`, email, password)
	// fmt.Println(payload)
	res, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if res.StatusCode != 200 {
		return "", errors.New(res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	out := TokenResponse{}
	if err := json.Unmarshal(body, &out); err != nil {
		return "", err
	}

	return out.AccessToken, nil
}

func saveToken(token string) error {

	return nil
}
