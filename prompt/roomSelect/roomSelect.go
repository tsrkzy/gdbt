package roomSelectPrompt

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	. "github.com/lepra-tsr/gdbt/config/room"
)

type RoomSelect struct {
}

func (u *RoomSelect) Ask(roomConfigJson *RoomConfigJson) error {
	fmt.Println("\ninput room's [id] and press return.\n")

	rooms := roomConfigJson.Rooms
	roomConfigJson.Show()

	fmt.Println("\nid?")
	buf := bufio.NewReader(os.Stdin)
	bytes, err := buf.ReadBytes('\n')
	if err != nil {
		return err
	}

	newRoomIndexStr := strings.TrimSpace(string(bytes))
	if newRoomIndexStr == "" {
		return nil
	}
	newRoomIndex, err := strconv.Atoi(newRoomIndexStr)
	if err != nil {
		return err
	}

	room := rooms[newRoomIndex]
	newRoomId := room.Id
	if err := roomConfigJson.SetCurrentById(newRoomId); err != nil {
		return err
	}
	roomConfigJson.Write()
	fmt.Println("you're in \"" + roomConfigJson.GetCurrentConnectedName() + "\" now.")
	fmt.Println("next, hit \"$ gdbt list\" to show messages.")
	return nil

	return errors.New("invalid room id:" + newRoomIndexStr)
}
