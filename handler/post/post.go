package post

import (
	"errors"
	"fmt"
	"github.com/lepra-tsr/gdbt/api/message"
	"github.com/lepra-tsr/gdbt/config/draft"
	"github.com/lepra-tsr/gdbt/config/room"
	"github.com/lepra-tsr/gdbt/prompt/confirm"
	"github.com/lepra-tsr/gdbt/vim"
	"regexp"
)

const (
	EditorMode     = "editor"
	DirectPostMode = "directPost"
	OpenDraftMode  = "openDraft"
	PostDraftMode  = "postDraft"
)

func Handler(inputStr string, mode string) error {
	// fmt.Println("post handler(" + mode + "): " + inputStr)

	switch mode {
	case EditorMode:
		if err := editorHandler(inputStr); err != nil {
			return err
		}
	case DirectPostMode:
		if err := directPostHandler(inputStr); err != nil {
			return err
		}
	case OpenDraftMode:
		if err := openDraftHandler(inputStr); err != nil {
			return err
		}
	case PostDraftMode:
		if err := postDraftHandler(inputStr); err != nil {
			return err
		}
	default:
		return errors.New("invalid key.")
	}

	return nil
}

func editorHandler(inputStr string) error {
	// テンポラリファイルを作成してvimで開く。
	// テンプレートを挿入。
	// 入力があったら追加する。
	// ファイル保存したら、テンポラリファイルを読み出す。
	//   保存せずに閉じたら終了。
	// 末尾の改行を取り、コメントを削除。(共通処理？)
	//   本文が空なら終了。
	// 本文を表示し、confirm。
	// enter ならば送信。
	//   e ならば再編集
	//   q ならば破棄して終了。
	//   dまたはそれ以外ならばドラフトを上書きして終了。
	roomJson := room.RoomConfigJson{}
	if err := roomJson.Read(); err != nil {
		return err
	}

	var err error
	currentRoomId := -1
	currentConnectedName := ""
	if roomId, err := roomJson.GetCurrentRoomId(); err != nil {
		fmt.Println("cannot find current room settings.")
		fmt.Println("please hit \"$ gdbt room\" to setup.")
		return err
	} else {
		currentRoomId = roomId
		currentConnectedName = roomJson.GetCurrentConnectedName()
	}

	vim := vim.Vim{}
	tempStr, err := vim.OpenTemporaryFile(inputStr)
	if err != nil {
		return err
	}

	text := clean(tempStr)

	fmt.Println("- - - - - - - ")
	fmt.Println(text)
	fmt.Println("- - - - - - - ")

	if text == "" {
		fmt.Println("empty lines.\nabort posting.")
		return nil
	}

	fmt.Println("post this message to \"" + currentConnectedName + "\"?")
	fmt.Println("y: Yes. post it.")
	fmt.Println("e: Edit(re-open).")
	fmt.Println("q: Quit. discard all texts.(not save)")
	fmt.Println("d: Draft. replace draft file with it.")
	fmt.Println("(y/e/q/d)?")

	confirm := confirmPrompt.Confirm{}
	command, err := confirm.AskIn("y,e,q,d")

	switch command {
	case "y":
		fmt.Println("")
		fmt.Println("... posting")
		postToRoom(text, currentRoomId)
		fmt.Println("post done.")
		return nil
	case "e":
		return editorHandler(text)
	case "q":
		fmt.Println("abort.")
		return nil
	case "d":
		/* overwrite draft. */
		fmt.Println("... save to draft")
		draftFile := draft.DraftFile{}
		draftFile.Body = text
		if err := draftFile.Write(); err != nil {
			return err
		}
		fmt.Println("saved.")
		return nil
	}

	return nil
}

func postToRoom(text string, roomId int) error {
	messageJson := message.MessagePostJson{}
	messageJson.RoomId = roomId
	messageJson.Source = text
	messageJson.Format = "markdown"
	if err := messageJson.Post(); err != nil {
		return err
	}
	return nil
}

func directPostHandler(inputStr string) error {
	fmt.Println(" -> directPostHandler.")
	// 末尾の改行を取り、コメントを削除。(共通処理？)
	//   本文が空ならエラー。
	// 本文を表示し、confirm。
	// enter ならば送信。
	//   e ならば再編集
	//   q ならば破棄して終了。
	//   dまたはそれ以外ならばドラフトを上書きして終了。
	return nil
}
func openDraftHandler(inputStr string) error {
	fmt.Println(" -> openDraftHandler.")
	// ドラフトファイルを開く。
	// vim側で保存/破棄して完結する想定。
	return nil
}
func postDraftHandler(inputStr string) error {
	fmt.Println(" -> postDraftHandler.")
	// ドラフトファイルを読み込む。
	// 末尾の改行を取り、コメントを削除。(共通処理？)
	//   本文が空ならエラー。
	// 本文を表示し、confirm。
	// enter ならば送信。
	//   e ならば再編集
	//   q ならば破棄して終了。
	//   d またはそれ以外ならばドラフトを上書きして終了。
	return nil
}

func clean(str string) string {
	reTrailEmptyLines := regexp.MustCompile(`(?m)(\s*)*\z`)
	replaced := reTrailEmptyLines.ReplaceAllString(str, "")
	return replaced
}
