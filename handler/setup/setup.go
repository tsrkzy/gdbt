package setup

// init は予約語なので避けた

import (
	"bufio"
	"bytes"
	"fmt"

	"encoding/json"
	"errors"
	"github.com/lepra-tsr/gdbt/util"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"syscall"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func Handler() error {
	fmt.Println("init handler.")
	// mkdir -p ~/.idbt
	// fopen("w", ~/.idbt/config)
	// fopen("w", ~/.idbt/draft)
	if err := checkConfigFileState(); err != nil {
		return err
	}

	// wait for input email and password
	// fetch accesstoken
	token, err := authPrompt()
	if err != nil {
		return err
	}
	fmt.Println("get access token done.")
	fmt.Println(token + " > ~/.idbt/config.json")

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

func checkConfigFileState() error {

	// mkdir -p ~/.idbt
	user, _ := user.Current()
	configFileDirName := ".idbt"
	configFileName := "config.json"
	draftFileName := "draft.md"
	configPath, _ := filepath.Abs(filepath.Join(user.HomeDir, configFileDirName))
	if err := os.Mkdir(configPath, 0774); err != nil {
		fmt.Println("mkdir: ~/" + configFileDirName + " is already exist.")
	}

	// open ~/.idbt/config (+write mode)
	configFilePath, _ := filepath.Abs(filepath.Join(user.HomeDir, configFileDirName, configFileName))
	if file, err := os.OpenFile(configFilePath, os.O_RDWR, 0774); err != nil {
		fmt.Println("cannot open: ~/" + configFilePath)
		return err
	} else {
		fmt.Println("file status check ok: " + configFilePath)
		file.Close()
	}

	// open ~/.idbt/draft (+write mode)
	draftFilePath, _ := filepath.Abs(filepath.Join(user.HomeDir, configFileDirName, draftFileName))
	if file, err := os.OpenFile(draftFilePath, os.O_RDWR, 0774); err != nil {
		fmt.Println("cannot open: ~/" + draftFilePath)
		return err
	} else {
		fmt.Println("file status check ok: " + draftFilePath)
		file.Close()
	}

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
		// fmt.Println("email: \"" + email + "\"")
	}

	if _password, err := askPassword(); err != nil {
		return "", err
	} else {
		password = _password
		// fmt.Println("password: " + util.IntToStr(len(password)) + "chars.")
	}

	if token, err := fetchToken(email, password); err != nil {
		return "", err
	} else {
		return token, nil
	}
}

func askEmail() (string, error) {
	fmt.Println("e-mail: ")
	buf := bufio.NewReader(os.Stdin)
	if email, err := buf.ReadBytes('\n'); err != nil {
		return "", err
	} else {
		return string(email), nil
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
			return string(bytePassword), nil
		}
		return "", err
	} else {
		return string(bytePassword), nil
	}
}

func fetchToken(email string, password string) (string, error) {
	// create json {grant_type, username,password}
	// curl
	em := util.StripNewLine(email)
	pw := util.StripNewLine(password)
	url := "https://idobata.io/oauth/token"
	payload := fmt.Sprintf(`{"grant_type":"password","username":"%v","password":"%v"}`, em, pw)

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
