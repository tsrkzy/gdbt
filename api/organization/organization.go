package organization

import (
	"encoding/json"

	"github.com/lepra-tsr/gdbt/api"
)

type OrganizationJson struct {
	Memberships   []api.Membership   `json:"memberships"`
	Organizations []api.Organization `json:"organizations"`
}

func (u *OrganizationJson) Fetch() error {
	bytes, err := api.CallGetWithCredential("/organizations")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, &u); err != nil {
		return err
	}

	return nil
}
