package credential

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/lepra-tsr/gdbt/config"
)

type CredentialJson struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func (u *CredentialJson) Write() error {
	file, err := os.OpenFile(CredentialJsonPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0774)
	if err != nil {
		return err
	} else {
		defer file.Close()
		jsonBytes, _ := json.Marshal(&u)
		file.Write(jsonBytes)
		
		return nil
	}
}

func (u *CredentialJson) Read() error {
	if bytes, err := ioutil.ReadFile(CredentialJsonPath); err != nil {
		return err
	} else {
		if err := json.Unmarshal(bytes, &u); err != nil {
			fmt.Println("Failed to parse credential.json")
			return err
		}
		return nil
	}
}
