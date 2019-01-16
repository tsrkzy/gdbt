package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (u *TokenResponse) Fetch(email string, password string) error {
	url := "https://idobata.io/oauth/token"
	payload := fmt.Sprintf(`{"grant_type":"password","username":"%v","password":"%v"}`, email, password)
	// fmt.Println(payload)
	res, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New(res.Status)
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, &u); err != nil {
		return err
	}

	return nil
}
