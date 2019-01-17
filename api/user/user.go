package user

import (
	"encoding/json"

	"github.com/lepra-tsr/gdbt/api"
)

type UserJson struct {
	Memberships []api.Membership `json:"memberships"`
	Joins       []api.Join       `json:"joins"`
	Users       []api.User       `json:"users"`
}

func (u *UserJson) Fetch() error {
	bytes, err := api.CallGetWithCredential("/users")
	if err != nil {
		return err
	}

	// fmt.Println(string(bytes))

	if err := json.Unmarshal(bytes, &u); err != nil {
		return err
	}

	return nil
}
