package room

import (
	"encoding/json"

	"github.com/lepra-tsr/gdbt/api"
)

type RoomJson struct {
	Joins    []api.Join `json:"joins"`
	BotJoins []api.Join `json:"bot_joins"`
	Rooms    []api.Room `json:"rooms"`
}

func (u *RoomJson) Fetch() error {
	bytes, err := api.CallGetWithCredential("/rooms")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, &u); err != nil {
		return err
	}

	return nil
}
