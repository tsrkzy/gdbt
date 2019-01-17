package message

import (
	"encoding/json"
	// "fmt"
	"github.com/lepra-tsr/gdbt/api"
	"github.com/lepra-tsr/gdbt/config/room"
)

type MessageJson struct {
	Messages []Message `json:"messages"`
	Meta     *Meta     `json:"meta"`
}

type Message struct {
	Id            int    `json:"id"`
	SenderName    string `json:"sender_name"`
	SenderIconUrl string `json:"sender_icon_url"`
	CreatedAt     string `json:"created_at"`
	Body          string `json:"body"`
	SenderId      int    `json:"sender_id"`
	SenderType    string `json:"sender_type"`
	RoomId        int    `json:"room_id"`
}

type Meta struct {
	HasNext bool `json:"has_next"`
}

func (u *MessageJson) Fetch(currentRoom *room.RoomInfo) error {
	messageUrl := currentRoom.MessageUrl
	bytes, err := api.CallGetWithCredential(messageUrl)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, &u); err != nil {
		return err
	}

	// fmt.Println(string(bytes))

	return nil
}
