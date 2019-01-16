package user

import "github.com/lepra-tsr/gdbt/api"
import "encoding/json"

type UserJson struct {
	Memberships []Membership `json:"memberships"`
	Joins       []Join       `json:"joins"`
	Users       []User       `json:"users"`
}

func (u *UserJson) Fetch() error {
	bytes, err := api.CallGetWithCredential("/users")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, &u); err != nil {
		return err
	}

	return nil
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
