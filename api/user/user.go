package user

import (
	"encoding/json"
	"fmt"

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

func (u *UserJson) Show() error {
	for i := 0; i < len(u.Users); i++ {
		user := u.Users[i]
		fmt.Println(user.Name, user.IconUrl)
	}

	return nil
}
