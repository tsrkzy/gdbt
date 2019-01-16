package api

import (
	"github.com/lepra-tsr/gdbt/config"
	"io/ioutil"
	"net/http"
)

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
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	IconUrl          string    `json:"icon_url"`
	Status           string    `json:"status"`
	Links            *UserLink `json:"links"`
	MembershipIdList []int     `json:"membership_ids"`
}

type UserLink struct {
	Stars string `json:"stars"`
}

type Organization struct {
	Id               int               `json:"id"`
	Name             string            `json:"name"`
	Slug             string            `json:"slug"`
	Links            *OrganizationLink `json:"links"`
	MembershipIdList []int             `json:"membership_ids"`
}

type OrganizationLink struct {
	Bots        string `json:"bots"`
	Invitations string `json:"invitations"`
	Rooms       string `json:"rooms"`
}

type Room struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	DescriptionSource string    `json:"description_source"`
	Description       string    `json:"description"`
	InvitationToken   string    `json:"invitation_token"`
	InvitationPolicy  string    `json:"invitation_policy"`
	Links             *RoomLink `json:"links"`
	OrganizationId    int       `json:"organization_id"`
	JoinIds           []int     `json:"join_ids"`
	BotJoinIds        []int     `json:"bot_join_ids"`
}

type RoomLink struct {
	HookEndpoints string `json:"hook_endpoints"`
	Messages      string `json:"messages"`
}

func CallGetWithCredential(path string) ([]byte, error) {
	_, token, err := config.ReadCredential()
	if err != nil {
		return nil, err
	}

	url := "https://idobata.io/api" + path
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-API-Token", token)
	req.Header.Set("User-Agent", "idbt")

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)

	return bytes, nil
}
