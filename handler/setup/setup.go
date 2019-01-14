package setup

// init は予約語なので避けた

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
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
