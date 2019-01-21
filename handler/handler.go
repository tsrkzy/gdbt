package handler

import (
	"regexp"

	"github.com/lepra-tsr/gdbt/api/message"
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
	reTrailEmptyLines := regexp.MustCompile(`(?m)(\s*)*\z`)
	replaced := reTrailEmptyLines.ReplaceAllString(str, "")
	return replaced
}
