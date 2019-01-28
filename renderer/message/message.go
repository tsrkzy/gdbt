package message

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	. "github.com/lepra-tsr/gdbt/api/message"
	"github.com/lepra-tsr/gdbt/config/room"
	"github.com/lepra-tsr/gdbt/util"
	"github.com/lunny/html2md"
)

type MessageRenderer struct {
	Messages []MessageInfo
}

type MessageInfo struct {
	Id         int
	Name       string
	CreatedAt  string
	Raw        string
	Markdown   string
	Replaced   string
	Compressed string
	LinkList   []Link
}

func compress(markdown string) string {
	replaced := markdown
	replaced = strings.Replace(replaced, "\n\n", "\n", -1)
	replaced = strings.Replace(replaced, "\n\n\n", "\n\n", -1)

	return replaced
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

		replacedString, linkList := getLinksFromMarkdown(info.Markdown)
		info.Replaced = replacedString
		info.Compressed = compress(info.Replaced)

		attachmentLinkList := getLinksFromAttachments(message.AttachImageUrlList, message.AttachFileUrlList)
		linkList = append(linkList, attachmentLinkList...)

		info.LinkList = linkList

		u.Messages = append(u.Messages, info)
	}
}

func strRFC3339toString(rfc3339str string) (string, error) {
	timestampUTC, err := time.Parse(time.RFC3339, rfc3339str)
	if err != nil {
		return "", err
	}
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	timestampJST := timestampUTC.In(jst)
	date := timestampJST.Format("Jan 02(Mon) 15:04")

	return date, nil
}

func (u *MessageRenderer) Show() {
	indentCount := 2
	indent := ""
	for i := 0; i < indentCount; i++ {
		indent = " " + indent
	}
	wallText := indent + "| "
	messages := u.Messages
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		date, err := strRFC3339toString(message.CreatedAt)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("\033[1m" + message.Name + "\033[0m" + " posted at [" + date + "]")
		lines := strings.Split(message.Compressed, "\n")
		for i := 0; i < len(lines); i++ {
			line := lines[i]
			fmt.Println(wallText + line)
		}
		for j := 0; j < len(message.LinkList); j++ {
			linkList := message.LinkList[j]
			fmt.Println(wallText + linkList.ToString())
		}
	}

	roomConfig := room.RoomConfigJson{}
	roomConfig.Read()
	roomName := roomConfig.GetCurrentConnectedName()
	fmt.Println(indent + "` " + roomName)
}

type Link struct {
	Label string
	Url   string
	Type  string
	Index int
}

func (u *Link) ToString() string {
	var label string
	if u.Label == "" {
		label = "*" + util.IntToStr(u.Index) + " " + u.Type
	} else {
		label = "*" + util.IntToStr(u.Index) + " " + u.Type + ": " + u.Label
	}
	str := label + " - " + u.Url
	return str
}

func getLinksFromMarkdown(compressedMarkdown string) (string, []Link) {
	replaced, inlineLinkList := findInlineLink(compressedMarkdown)
	linkList := []Link{}
	linkList = append(linkList, inlineLinkList...)

	return replaced, linkList
}

func getLinksFromAttachments(imageUrlList []string, fileUrlList []string) []Link {
	linkList := []Link{}
	for i := 0; i < len(imageUrlList); i++ {
		url := imageUrlList[i]
		link := Link{}
		link.Label = util.IntToStr(i)
		link.Url = url
		link.Type = "ATTACHED IMAGE"
		linkList = append(linkList, link)
	}

	for i := 0; i < len(fileUrlList); i++ {
		url := fileUrlList[i]
		duplicated := false
		for j := 0; j < len(imageUrlList); j++ {
			imgUrl := imageUrlList[j]
			duplicated = (url == imgUrl)
			if duplicated {
				break
			}
		}
		if duplicated {
			continue
		}

		link := Link{}
		link.Label = util.IntToStr(i)
		link.Url = url
		link.Type = "ATTACHED FILE"
		linkList = append(linkList, link)
	}

	return linkList
}

func findInlineLink(markdown string) (string, []Link) {
	linkList := []Link{}

	loops := 0
	const LoopLimit = 100
	str := markdown
	for loops < LoopLimit {
		loops++
		lines := strings.Split(str, "\n")
		replaced := false
		finished := false
		for i := 0; i < len(lines); i++ {
			line := lines[i]
			if replacedText, link, hitCount := findImage(line, loops); hitCount != 0 {
				lines[i] = replacedText
				str = strings.Join(lines, "\n")
				linkList = append(linkList, link)
				replaced = true
				finished = (len(lines) == (i + 1))
				break
			}

			if replacedText, link, hitCount := findAnchor(line, loops); hitCount != 0 {
				lines[i] = replacedText
				str = strings.Join(lines, "\n")
				linkList = append(linkList, link)
				replaced = true
				finished = (len(lines) == (i + 1))
				break
			}

			if replaced {
				continue
			}
		}

		if finished {
			break
		}
	}

	// fmt.Println(linkList)
	return str, linkList
}

func findImage(str string, index int) (string, Link, int) {
	urlPattern := `https?:\/\/[a-zA-Z0-9\-\_\.\!\'\*\;\/\?\:\@\&\=\+\$\%\#\,]+`
	WithoutTitlePattern := `(?m)\!\[()image\]\((` + urlPattern + `)\)`
	reWithoutTitle := regexp.MustCompile(WithoutTitlePattern)

	WithTitlePattern := `(?m)\!\[([^\]]+)\]\((` + urlPattern + `)\)`
	reWithTitle := regexp.MustCompile(WithTitlePattern)

	Pattern := `(?m)\!\[()[^\]]*\]\((` + urlPattern + `)\)`
	re := regexp.MustCompile(Pattern)

	if replacedText, link, hitCount := findImageByPattern(reWithoutTitle, str, index); hitCount != 0 {
		return replacedText, link, hitCount
	}

	if replacedText, link, hitCount := findImageByPattern(reWithTitle, str, index); hitCount != 0 {
		return replacedText, link, hitCount
	}

	if replacedText, link, hitCount := findImageByPattern(re, str, index); hitCount != 0 {
		return replacedText, link, hitCount
	}

	return str, Link{}, 0
}

func findImageByPattern(pattern *regexp.Regexp, str string, index int) (string, Link, int) {
	result := pattern.FindStringSubmatch(str)
	if len(result) == 0 {
		/* not found */
		return str, Link{}, 0
	}

	matchedText := result[0]
	label := result[1]
	if label == "" {
		label = "IMAGE"
	}
	url := result[2]

	link := Link{}
	link.Type = "IMAGE"
	link.Label = label
	link.Url = url
	link.Index = index

	/* replace */
	altText := "![" + label + "](*" + util.IntToStr(index) + ")"
	str = strings.Replace(str, matchedText, altText, -1)

	return str, link, len(result)
}

func findAnchor(str string, index int) (string, Link, int) {
	urlPattern := `https?:\/\/[a-zA-Z0-9\-\_\.\!\'\*\;\/\?\:\@\&\=\+\$\%\#\,]+`
	anchorPattern := `(?m)\[([^\]]+)\]\((` + urlPattern + `)\)`
	reAnchor := regexp.MustCompile(anchorPattern)

	if replacedText, link, hitCount := findAnchorByPattern(reAnchor, str, index); hitCount != 0 {
		return replacedText, link, hitCount
	}

	return str, Link{}, 0
}

func findAnchorByPattern(pattern *regexp.Regexp, str string, index int) (string, Link, int) {
	result := pattern.FindStringSubmatch(str)
	if len(result) == 0 {
		/* not found */
		return str, Link{}, 0
	}

	matchedText := result[0]
	label := result[1]
	url := result[2]

	link := Link{}
	link.Type = "LINK"
	link.Label = label
	link.Index = index
	if label == url {
		link.Label = "link"
	}
	link.Url = url

	/* replace */
	altText := "[" + link.Label + "](*" + util.IntToStr(index) + ")"
	str = strings.Replace(str, matchedText, altText, -1)

	return str, link, len(result)
}

func (u *MessageInfo) countRows() int {
	/* 本文の行数 */
	re := regexp.MustCompile(`(?m)$`)
	result := re.FindAllStringSubmatch(u.Compressed, -1)
	contentRowCount := len(result)

	/* 本文中に挿入したurl */
	inlineLinkCount := len(u.LinkList)

	return contentRowCount + inlineLinkCount
}
