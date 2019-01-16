package setup

// init は予約語なので避けた

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/lepra-tsr/gdbt/api/user"
	"github.com/lepra-tsr/gdbt/config"
	"github.com/lepra-tsr/gdbt/prompt/auth"
	"io/ioutil"
	"net/http"
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
	_, err := startAuthPrompt()
	if err != nil {
		return err
	}
	fmt.Println("get access token done.")

	userJson := UserJson{}
	if err := userJson.Fetch(); err != nil {
		return err
	}

	fmt.Println(userJson.Users[0])

	// if err := fetchChannelEntity(); err != nil {
	// 	return err
	// }
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

	if token, err := fetchToken(email, password); err != nil {
		return "", err
	} else {
		if err := config.WriteCredential(email, token); err != nil {
			return "", err
		}

		return token, nil
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

func fetchChannelEntity() error {
	url := "https://idobata.io/api/users"
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))
	// payload :=
	return nil
}
