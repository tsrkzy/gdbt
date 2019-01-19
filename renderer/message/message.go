package message

import (
	"fmt"
	"regexp"

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
		fmt.Println(info.Markdown)
		fmt.Println("---")
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
}

func (u *Link) ToString() string {
	var label string
	if u.Label == "" {
		label = u.Type
	} else {
		label = u.Type + ": " + u.Label
	}
	str := "[" + label + "] - " + u.Url
	return str
}

func getLinksFromMarkdown(compressedMarkdown string) (string, []Link) {
	// ![***](uri)、[***](uri)のパターンでマッチ
	// [IMAGE n:***]、[LINK n:***]に置換
	// 以下の行を追加
	// [type: ***] - uri
	imageReplaced, inlineImageList := findInlineLink(compressedMarkdown)
	replaced, inlineLinkList := findInlineLink(imageReplaced)
	linkList := []Link{}
	linkList = append(linkList, inlineImageList...)
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
	// 変更予定
	// 画像要素の抽出、置換
	// ループ開始
	//   画像要素でパターンマッチし、該当する文字列を取得
	//   該当する文字列をすべて置換

	linkList := []Link{}
	urlPattern := `https?:\/\/[a-zA-Z0-9\-\_\.\!\'\(\)\*\;\/\?\:\@\&\=\+\$\%\#\,]+`

	imageWithoutTitlePattern := `(?m)\!\[()image\]\((` + urlPattern + `)\)`
	reImageWithoutTitle := regexp.MustCompile(imageWithoutTitlePattern)
	resultImageWithoutTitleList := reImageWithoutTitle.FindAllStringSubmatch(markdown, -1)
	replaced := reImageWithoutTitle.ReplaceAllString(markdown, "[IMAGE]")

	imageWithTitlePattern := `(?m)\!\[([^\]]+)\]\((` + urlPattern + `)\)`
	reImageWithTitle := regexp.MustCompile(imageWithTitlePattern)
	resultImageWithTitleList := reImageWithTitle.FindAllStringSubmatch(replaced, -1)
	replaced = reImageWithTitle.ReplaceAllString(replaced, "[IMAGE: $1]")

	imagePattern := `(?m)\!\[()[^\]]*\]\((` + urlPattern + `)\)`
	reImage := regexp.MustCompile(imagePattern)
	resultImageList := reImage.FindAllStringSubmatch(replaced, -1)
	replaced = reImage.ReplaceAllString(replaced, "[IMAGE]")

	resultImageList = append(resultImageList, resultImageWithoutTitleList...)
	resultImageList = append(resultImageList, resultImageWithTitleList...)
	for i := 0; i < len(resultImageList); i++ {
		result := resultImageList[i]
		label := result[1]
		url := result[2]
		link := Link{}
		link.Type = "INLINE IMAGE"
		link.Label = label
		link.Url = url
		linkList = append(linkList, link)
	}

	nastyAnchorPattern := `(?m)\[(` + urlPattern + `)\]\((` + urlPattern + `)\)`
	reNastyAnchor := regexp.MustCompile(nastyAnchorPattern)
	resultNastyAnchorList := reNastyAnchor.FindAllStringSubmatch(replaced, -1)
	replaced = reNastyAnchor.ReplaceAllString(replaced, "[ANCHOR]")

	anchorPattern := `(?m)\[([^\]]+)\]\((` + urlPattern + `)\)`
	reAnchor := regexp.MustCompile(anchorPattern)
	resultAnchorList := reAnchor.FindAllStringSubmatch(replaced, -1)
	replaced = reAnchor.ReplaceAllString(replaced, "[ANCHOR: $1]")

	resultAnchorList = append(resultAnchorList, resultNastyAnchorList...)
	for i := 0; i < len(resultAnchorList); i++ {
		result := resultAnchorList[i]
		label := result[1]
		url := result[2]

		if label == url {
			label = ""
		}
		link := Link{}
		link.Type = "ANCHOR"
		link.Label = label
		link.Url = url
		linkList = append(linkList, link)
	}

	return replaced, linkList
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
