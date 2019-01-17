package roomSelectPrompt

import (
	"bufio"
	"errors"
	"fmt"
	. "github.com/lepra-tsr/gdbt/config/room"
	"github.com/lepra-tsr/gdbt/util"
	"os"
	"strings"
)

type RoomSelect struct {
}

func (u *RoomSelect) Ask(roomConfigJson *RoomConfigJson) error {
	fmt.Println("\ninput room's [id] and press return.\n")

	rooms := roomConfigJson.Rooms
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		id := util.IntToStr(room.Id)
		fmt.Println("[" + id + "] " + room.GetConnectedName())
	}
	fmt.Println("\nid?")
	buf := bufio.NewReader(os.Stdin)
	bytes, err := buf.ReadBytes('\n')
	if err != nil {
		return err
	}

	newRoomId := strings.TrimSpace(string(bytes))
	if newRoomId == "" {
		return nil
	}

	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		id := util.IntToStr(room.Id)
		if id != newRoomId {
			continue
		}
		if err := roomConfigJson.SetCurrentById(newRoomId); err != nil {
			return err
		}
		roomConfigJson.Write()
		fmt.Println("you're in \"" + roomConfigJson.GetCurrentConnectedName() + "\" now.")
		return nil
	}

	return errors.New("invalid room id:" + newRoomId)
}
