package user

import "github.com/lepra-tsr/gdbt/api"
import "encoding/json"

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

	if err := json.Unmarshal(bytes, &u); err != nil {
		return err
	}

	return nil
}
