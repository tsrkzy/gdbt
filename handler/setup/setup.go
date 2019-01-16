package setup

// init は予約語なので避けた

import (
	"fmt"
	. "github.com/lepra-tsr/gdbt/api/token"
	. "github.com/lepra-tsr/gdbt/api/user"
	. "github.com/lepra-tsr/gdbt/api/organization"
	. "github.com/lepra-tsr/gdbt/api/room"
	"github.com/lepra-tsr/gdbt/config"
	"github.com/lepra-tsr/gdbt/prompt/auth"
)

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

	// _, err := startAuthPrompt()
	// if err != nil {
	// 	return err
	// }

	userJson := UserJson{}
	if err := userJson.Fetch(); err != nil {
		return err
	}

	organizationJson := OrganizationJson{}
	if err := organizationJson.Fetch(); err != nil {
		return err
	}

	roomJson := RoomJson{}
	if err := roomJson.Fetch(); err != nil {
		return err
	}

	fmt.Println(roomJson.Rooms[0])

	// write userName
	// write accessToken
	// write user
	// write organizations
	// write rooms

	// wait for input room selection

	return nil
}

func startAuthPrompt() (string, error) {
	var (
		email    string
		password string
	)
	if _email, err := authPrompt.AskEmail(); err != nil {
		return "", err
	} else {
		email = _email
	}

	if _password, err := authPrompt.AskPassword(); err != nil {
		return "", err
	} else {
		password = _password
	}

	tokenResponse := TokenResponse{}
	if err := tokenResponse.Fetch(email, password); err != nil {
		return "", err
	}

	token := tokenResponse.AccessToken
	if err := config.WriteCredential(email, token); err != nil {
		return "", err
	}

	return token, nil
}
