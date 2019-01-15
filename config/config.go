package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

type CredentialJson struct {
	UserName string `json:"username"`
	Token    string `json:"token"`
}

type Channels struct {
	List []Channel `json:"list"`
}

type Channel struct {
	id               int
	name             string
	messageUrl       string
	organizationName string
}

const (
	configFileDirName  = ".gdbt"
	credentialJsonName = "credential.json"
	channelJsonName    = "channels.json"
	draftFileName      = "draft.md"
)

var (
	ConfigPath         string
	CredentialJsonPath string
	ChannelJsonPath    string
	DraftFilePath      string
)

func init() {
	user, _ := user.Current()

	ConfigPath, _ = filepath.Abs(filepath.Join(user.HomeDir, configFileDirName))
	CredentialJsonPath, _ = filepath.Abs(filepath.Join(user.HomeDir, configFileDirName, credentialJsonName))
	ChannelJsonPath, _ = filepath.Abs(filepath.Join(user.HomeDir, configFileDirName, channelJsonName))
	DraftFilePath, _ = filepath.Abs(filepath.Join(user.HomeDir, configFileDirName, draftFileName))
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func CheckConfigFileState() error {

	// mkdir -p ~/.gdbt
	if err := os.MkdirAll(ConfigPath, 0774); err != nil {
		return err
	}

	// open ~/.gdbt/config (+write mode)
	if err := OpenOrCreateFileWithWriteMode(CredentialJsonPath); err != nil {
		return err
	}

	if err := OpenOrCreateFileWithWriteMode(ChannelJsonPath); err != nil {
		return err
	}

	// open ~/.gdbt/draft (+write mode)
	if err := OpenOrCreateFileWithWriteMode(DraftFilePath); err != nil {
		return err
	}

	return nil
}

func OpenOrCreateFileWithWriteMode(path string) error {
	if Exists(path) {
		// 存在する場合は書き込みモードで開けるかチェック
		if file, err := os.OpenFile(path, os.O_RDWR, 0774); err != nil {
			fmt.Println("cannot open: ~/" + path)
			return err
		} else {
			defer file.Close()
			fmt.Println("file status check ok: " + path)
		}
	} else {
		// 存在しない場合は作成
		if file, err := os.Create(path); err != nil {
			fmt.Println("cannot create: ~/" + path)
			return err
		} else {
			defer file.Close()
			fmt.Println("file status check ok: " + path)
		}
	}

	return nil
}

func WriteCredential(username string, token string) error {
	if file, err := os.OpenFile(CredentialJsonPath, os.O_RDWR, 0774); err != nil {
		return err
	} else {
		defer file.Close()
		jsonStr := fmt.Sprintf(`{"username":"%v","token":"%v"}`, username, token)
		fmt.Fprintln(file, jsonStr)
		return nil
	}
}

func ReadCredential() (string, string, error) {
	if bytes, err := ioutil.ReadFile(CredentialJsonPath); err != nil {
		return "", "", err
	} else {
		credentialJson := CredentialJson{}
		if err := json.Unmarshal(bytes, &credentialJson); err != nil {
			return "", "", err
		}

		return credentialJson.UserName, credentialJson.Token, nil
	}
}

func WriteChannels(channelList *[]Channel) error {

	return nil
}
