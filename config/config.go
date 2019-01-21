package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

const (
	configFileDirName  = ".gdbt"
	credentialJsonName = "credential.json"
	roomJsonName       = "rooms.json"
	draftFileName      = "draft.md"
)

var (
	ConfigPath         string
	CredentialJsonPath string
	RoomJsonPath       string
	DraftFilePath      string
)

func init() {
	user, _ := user.Current()

	ConfigPath, _ = filepath.Abs(filepath.Join(user.HomeDir, configFileDirName))
	CredentialJsonPath, _ = filepath.Abs(filepath.Join(user.HomeDir, configFileDirName, credentialJsonName))
	RoomJsonPath, _ = filepath.Abs(filepath.Join(user.HomeDir, configFileDirName, roomJsonName))
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

	if err := OpenOrCreateFileWithWriteMode(RoomJsonPath); err != nil {
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
