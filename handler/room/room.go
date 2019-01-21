package room

import (
	"fmt"

	. "github.com/lepra-tsr/gdbt/config/room"
	"github.com/lepra-tsr/gdbt/handler/setup"
	. "github.com/lepra-tsr/gdbt/prompt/roomSelect"
)

const (
	SelectRoomMode  = "selectRoom"
	ReloadAndSelect = "reloadAndSelect"
	Show            = "show"
)

func Handler(reloadFlag bool, showFlag bool) error {

	mode := SelectRoomMode
	if reloadFlag {
		mode = ReloadAndSelect
	} else if showFlag {
		mode = Show
	} else {
		mode = SelectRoomMode
	}

	switch mode {
	case SelectRoomMode:
		return SelectRoomHandler()
	case ReloadAndSelect:
		return ReloadAndSelectHandler()
	case Show:
		return ShowHandler()
	}

	return nil
}

func SelectRoomHandler() error {
	roomConfigJson := RoomConfigJson{}
	roomConfigJson.Read()
	roomSelect := RoomSelect{}
	if err := roomSelect.Ask(&roomConfigJson); err != nil {
		fmt.Println(err)
	}
	return nil
}

func ReloadAndSelectHandler() error {
	fmt.Println("fetching room infomation...")
	if err := setup.UpdateRoomConfig(); err != nil {
		return err
	}
	fmt.Println("done.")
	return SelectRoomHandler()
}

func ShowHandler() error {
	roomConfig := RoomConfigJson{}
	roomConfig.Read()
	roomConfig.Show()
	return nil
}
