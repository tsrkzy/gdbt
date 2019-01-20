package message

import (
	"fmt"
	"regexp"
	"strings"

	. "github.com/lepra-tsr/gdbt/api/message"
	"github.com/lepra-tsr/gdbt/util"

	// "github.com/lepra-tsr/gdbt/util"
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

		replacedString, linkList := getLinksFromMarkdown(info.Markdown)
		info.Replaced = replacedString

		attachmentLinkList := getLinksFromAttachments(message.AttachImageUrlList, message.AttachFileUrlList)
		linkList = append(linkList, attachmentLinkList...)

		info.LinkList = linkList

		fmt.Println("===")
		// fmt.Println(info.Markdown)
		// fmt.Println("---")
		fmt.Println(info.Replaced)
		fmt.Println("---")
		for i := 0; i < len(info.LinkList); i++ {
			linkList := info.LinkList[i]
			fmt.Println(linkList.ToString())
		}

		u.Messages = append(u.Messages, info)
	}
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
	url := result[2]

	link := Link{}
	link.Type = "IMAGE"
	link.Label = label
	link.Url = url
	link.Index = index

	/* replace */
	altText := "IMAGE(*" + util.IntToStr(index) + ")"
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
		link.Label = ""
	}
	link.Url = url

	/* replace */
	altText := "LINK(*" + util.IntToStr(index) + ")"
	str = strings.Replace(str, matchedText, altText, -1)

	return str, link, len(result)
}

func (u *MessageRenderer) Show() {
	messages := u.Messages
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		message.countRows()
	}
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

func (u *MessageInfo) countRows() int {
	/* 本文の行数 */
	re := regexp.MustCompile(`(?m)$`)
	result := re.FindAllStringSubmatch(u.Compressed, -1)
	contentRowCount := len(result)

	/* 本文中に挿入したurl */
	inlineLinkCount := len(u.LinkList)

	// fmt.Println(u.Compressed)
	// fmt.Println("contentRowCount: " + util.IntToStr(contentRowCount) + ", " + "attachImageCount: " + util.IntToStr(attachImageCount) + ", " + "attachFileCount: " + util.IntToStr(attachFileCount) + ", " + "inlineLinkCount: " + util.IntToStr(inlineLinkCount) + ", ")
	// fmt.Println("===")
	return contentRowCount /* + attachImageCount + attachFileCount */ + inlineLinkCount
}
