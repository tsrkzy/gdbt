package list

import (
	"errors"

	"github.com/lepra-tsr/gdbt/api/message"
	. "github.com/lepra-tsr/gdbt/config/room"
	. "github.com/lepra-tsr/gdbt/renderer/message"
)

func Handler(unionFlag bool) error {
	if unionFlag {
		return unionHandler()
	} else {
		return currentHandler()
	}
}

func currentHandler() error {
	roomConfigJson := RoomConfigJson{}
	roomConfigJson.Read()
	currentRoom := roomConfigJson.CurrentRoom

	if currentRoom == nil {
		return errors.New("please select room with command following.\n $ gdbt room")
	}

	messageJson := message.MessageJson{}
	err := messageJson.Fetch(currentRoom)
	if err != nil {
		return err
	}

	messageRenderer := MessageRenderer{}
	messageRenderer.ParseMessageJson(&messageJson)

	// fmt.Println(messageRenderer)
	messageRenderer.Show("current")

	return nil
}

func unionHandler() error {
	messageJson := message.MessageJson{}
	err := messageJson.Fetch(nil)
	if err != nil {
		return err
	}

	messageRenderer := MessageRenderer{}
	messageRenderer.ParseMessageJson(&messageJson)

	messageRenderer.Show("union")

	return nil
}
