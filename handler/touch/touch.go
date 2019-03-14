package touch

// init は予約語なので避けた

import (
	"fmt"
	"github.com/lepra-tsr/gdbt/config/room"
)

func Handler() error {
	fmt.Println("start touch.")
	roomConfig := room.RoomConfigJson{}
	roomConfig.Read()
	roomConfig.Touch()
	return nil
}
