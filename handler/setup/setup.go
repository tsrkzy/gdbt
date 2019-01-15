package setup

// init は予約語なので避けた

import (
	"bufio"
	"bytes"
	"fmt"

	// "io/ioutil"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

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
	if err := authPrompt(); err != nil {
		return err
	}

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

func authPrompt() error {
	var (
		email    string
		password string
	)
	if _email, err := askEmail(); err != nil {
		return err
	} else {
		email = _email
		fmt.Println(email)
	}

	if _password, err := askPassword(); err != nil {
		return err
	} else {
		password = _password
		fmt.Println(password)
	}

	if token, err := fetchToken(email, password); err != nil {
		return err
	} else {
		fmt.Println(token)

	}

	return nil
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
		return "", err
	} else {
		return string(bytePassword), nil
	}
}

func fetchToken(email string, password string) (string, error) {
	// create json {grant_type, username,password}
	// curl
	url := "https://idobata.io/oauth/token"
	payload := "{\"grant_type\":\"password\",\"username\":\"tsrmix@gmail.com\",\"password\":\"#xatm0920\"}"

	req, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(req)

	return "auth token", nil
}
