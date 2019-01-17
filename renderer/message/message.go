package message

import (
	"fmt"
	"regexp"

	. "github.com/lepra-tsr/gdbt/api/message"
	"github.com/lepra-tsr/gdbt/util"
	"github.com/lunny/html2md"
)

type MessageRenderer struct {
	Messages []MessageInfo
}

func compress(markdown string) string {
	return markdown
}

func (u *MessageRenderer) ParseMessageJson(json *MessageJson) {
	messages := json.Messages
	for i := 0; i < len(messages); i++ {
		info := MessageInfo{}
		message := messages[i]
		info.Id = message.Id
		info.Name = message.SenderName
		info.CreatedAt = message.CreatedAt
		info.Raw = message.Body
		info.Markdown = html2md.Convert(message.Body)
		info.Compressed = compress(info.Markdown)
		info.AttachedImages = append(info.AttachedImages, message.AttachImageUrlList...)
		info.AttachedFiles = append(info.AttachedFiles, message.AttachFileUrlList...)
		u.Messages = append(u.Messages, info)
	}
}

func (u *MessageRenderer) Show() {
	messages := u.Messages
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		message.countRows()
	}
}

type MessageInfo struct {
	Id             int
	Name           string
	CreatedAt      string
	Raw            string
	Markdown       string
	Compressed     string
	AttachedImages []string
	AttachedFiles  []string
	InlineUrlList  []string
}

func (u *MessageInfo) countRows() int {
	/* 本文の行数 */
	re := regexp.MustCompile(`(?m)$`)
	result := re.FindAllStringSubmatch(u.Compressed, -1)
	contentRowCount := len(result)

	/* 添付 */
	attachImageCount := len(u.AttachedImages)
	attachFileCount := len(u.AttachedFiles)

	/* 本文中に挿入したurl */
	inlineLinkCount := len(u.InlineUrlList)

	fmt.Println(u.Compressed)
	fmt.Println("contentRowCount: " + util.IntToStr(contentRowCount) + ", " + "attachImageCount: " + util.IntToStr(attachImageCount) + ", " + "attachFileCount: " + util.IntToStr(attachFileCount) + ", " + "inlineLinkCount: " + util.IntToStr(inlineLinkCount) + ", ")
	fmt.Println("===")
	return contentRowCount + attachImageCount + attachFileCount + inlineLinkCount
}
