package setup

// init は予約語なので避けた

import (
	"fmt"

	. "github.com/lepra-tsr/gdbt/api/organization"
	. "github.com/lepra-tsr/gdbt/api/room"
	. "github.com/lepra-tsr/gdbt/api/token"
	. "github.com/lepra-tsr/gdbt/api/user"
	"github.com/lepra-tsr/gdbt/config"
	. "github.com/lepra-tsr/gdbt/config/channel"
	. "github.com/lepra-tsr/gdbt/config/credential"
	authPrompt "github.com/lepra-tsr/gdbt/prompt/auth"
)

func Handler() error {
	fmt.Println("init handler.")
	if err := config.CheckConfigFileState(); err != nil {
		return err
	}
	/* 開発時にemailとtokenを更新しない場合はここをコメントアウト */
	// if err := writeCredential(); err != nil {
	// 	return err
	// }

	if err := writeChannel(); err != nil {
		return err
	}
	// write user
	// write organizations
	// write rooms

	// wait for input room selection

	return nil
}

func writeCredential() error {
	auth := authPrompt.Auth{}
	if err := auth.AskEmail(); err != nil {
		return err
	}
	if err := auth.AskPassword(); err != nil {
		return err
	}
	email := auth.Email
	password := auth.Password

	tokenResponse := TokenResponse{}
	if err := tokenResponse.Fetch(email, password); err != nil {
		return err
	}

	token := tokenResponse.AccessToken
	credential := CredentialJson{}
	credential.Token = token
	credential.Email = email
	if err := credential.Write(); err != nil {
		return err
	}
	return nil
}

func writeChannel() error {

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

	configJson := ConfigJson{}
	configJson.ParseServerEntity(&userJson, &organizationJson, &roomJson)
	configJson.Write()

	return nil
}
