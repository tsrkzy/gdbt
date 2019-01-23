package handler

import (
	"github.com/lepra-tsr/gdbt/api/message"
	"github.com/lepra-tsr/gdbt/util"
)

func PostToRoom(text string, roomId int) error {
	messageJson := message.MessagePostJson{}
	messageJson.RoomId = roomId
	messageJson.Source = text
	messageJson.Format = "markdown"
	if err := messageJson.Post(); err != nil {
		return err
	}
	return nil
}

func Clean(str string) string {
	replaced := util.RemoveCommentLines(str)
	replaced = util.RemoveTrailEmptyLines(replaced)
	return replaced
}
