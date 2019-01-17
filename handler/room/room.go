package room

import (
	"fmt"
	. "github.com/lepra-tsr/gdbt/config/room"
	. "github.com/lepra-tsr/gdbt/prompt/roomSelect"
)

func Handler() error {
	roomConfigJson := RoomConfigJson{}
	roomConfigJson.Read()
	roomSelect := RoomSelect{}
	if err := roomSelect.Ask(&roomConfigJson); err != nil{
		fmt.Println(err)
	}
	return nil
}
