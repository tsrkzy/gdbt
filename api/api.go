package api

import (
	"io/ioutil"
	"net/http"

	. "github.com/lepra-tsr/gdbt/config/credential"
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
	Id                       int       `json:"id"`
	Name                     string    `json:"name"`
	IconUrl                  string    `json:"icon_url"`
	Status                   string    `json:"status"`
	Email                    string    `json:"email"`
	ChannelName              string    `json:"channel_name"`
	CreatedAt                string    `json:"created_at"`
	EnterBehaviorDesktop     string    `json:"enter_behavior_desktop"`
	EnterBehaviorMobile      string    `json:"enter_behavior_mobile"`
	UseMarkdown              bool      `json:"use_markdown"`
	MessageFoldable          bool      `json:"message_foldable"`
	ReceiveBroadcastMentions bool      `json:"receive_broadcast_mentions"`
	Links                    *UserLink `json:"links"`
	MembershipIdList         []int     `json:"membership_ids"`
	JoinIdList               []int     `json:"join_ids"`
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
	credential := CredentialJson{}
	if err := credential.Read(); err != nil {
		return nil, err
	}
	token := credential.Token

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

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
