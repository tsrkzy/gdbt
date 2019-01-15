package setup

// init は予約語なので避けた

import (
	"bufio"
	"bytes"
	"fmt"

	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"syscall"

	"github.com/lepra-tsr/gdbt/config"
	"github.com/lepra-tsr/gdbt/util"
	"golang.org/x/crypto/ssh/terminal"
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
	_, err := authPrompt()
	if err != nil {
		return err
	}
	fmt.Println("get access token done.")

	if _, err := callGetWithCredential("/users"); err != nil {
		return err
	}

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

func authPrompt() (string, error) {
	var (
		email    string
		password string
	)
	if _email, err := askEmail(); err != nil {
		return "", err
	} else {
		email = _email
	}

	if _password, err := askPassword(); err != nil {
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

func askEmail() (string, error) {
	fmt.Println("e-mail: ")
	buf := bufio.NewReader(os.Stdin)
	if email, err := buf.ReadBytes('\n'); err != nil {
		return "", err
	} else {
		return util.StripNewLine(string(email)), nil
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
			return util.StripNewLine(string(bytePassword)), nil
		}
	} else {
		return util.StripNewLine(string(bytePassword)), nil
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

type UserResponseJson struct {
	Memberships []Membership `json:"memberships"`
	Joins       []Join       `json:"joins"`
	Users       []User       `json:"users"`
}

type Membership struct {
	Id             int    `json:"id"`
	Role           string `json:"role"`
	OrganizationId int    `json:"organization_id"`
	GuyId          int    `json:"guy_id"`
}

type Join struct {
	Id     int `json:"id"`
	RoomId int `json:"room_id"`
	GuyId  int `json:"guy_id"`
}

type User struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	IconUrl          string `json:"icon_url"`
	Status           string `json:"status"`
	Links            *Link  `json:"links"`
	MembershipIdList []int  `json:"membership_ids"`
}

type Link struct {
	Stars string `json:"stars"`
}

func callGetWithCredential(path string) (string, error) {
	_, token, err := config.ReadCredential()
	if err != nil {
		return "", err
	}

	url := "https://idobata.io/api" + path
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-API-Token", token)
	req.Header.Set("User-Agent", "idbt")

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)

	userResponseJson := UserResponseJson{}
	if err := json.Unmarshal(bytes, &userResponseJson); err != nil {
		return "", err
	}

	fmt.Println(userResponseJson.Memberships[0])
	return "", nil
}
