package post

import (
	"errors"
	"fmt"

	"github.com/lepra-tsr/gdbt/api/message"
	"github.com/lepra-tsr/gdbt/config/draft"
	"github.com/lepra-tsr/gdbt/config/room"
	"github.com/lepra-tsr/gdbt/handler"
	confirmPrompt "github.com/lepra-tsr/gdbt/prompt/confirm"
	"github.com/lepra-tsr/gdbt/vim"
)

const (
	EditorMode     = "editor"
	DirectPostMode = "directPost"
	OpenDraftMode  = "openDraft"
	PostDraftMode  = "postDraft"
)

func Handler(messageOption string, draftFlag bool) error {
	// fmt.Println("post handler(" + mode + "): " + inputStr)

	mode := EditorMode
	if messageOption != "" {
		mode = DirectPostMode
	} else if draftFlag {
		mode = PostDraftMode
	} else {
		mode = EditorMode
	}

	switch mode {
	case EditorMode:
		if err := editorHandler(messageOption); err != nil {
			return err
		}
	case DirectPostMode:
		if err := directPostHandler(messageOption); err != nil {
			return err
		}
	case PostDraftMode:
		if err := postDraftHandler(); err != nil {
			return err
		}
	default:
		return errors.New("invalid key.")
	}

	return nil
}
func getCurrentRoom() (*room.RoomInfo, error) {
	roomJson := room.RoomConfigJson{}
	if err := roomJson.Read(); err != nil {
		return nil, err
	}

	if roomJson.CurrentRoom == nil {
		errMsg := "cannot find current room settings. please hit \"$ gdbt room\" to setup."
		return nil, errors.New(errMsg)
	}

	return roomJson.CurrentRoom, nil
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
	roomInfo, err := getCurrentRoom()
	if err != nil {
		return err
	}

	vim := vim.Vim{}
	tempStr, err := vim.OpenTemporaryFile(inputStr)
	if err != nil {
		return err
	}

	text := handler.Clean(tempStr)

	return confirmBeforePost(roomInfo, text)
}

func confirmBeforePost(roomInfo *room.RoomInfo, text string) error {
	currentRoomId := roomInfo.Id
	currentConnectedName := roomInfo.GetConnectedName()

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
	if err != nil {
		return err
	}

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
	default:
		return errors.New("invalid subcommand: \"" + command)
	}
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
	// 末尾の改行を取り、コメントを削除。(共通処理？)
	//   本文が空ならエラー。
	// 本文を表示し、confirm。
	// enter ならば送信。
	//   e ならば再編集
	//   q ならば破棄して終了。
	//   dまたはそれ以外ならばドラフトを上書きして終了。
	roomInfo, err := getCurrentRoom()
	if err != nil {
		return err
	}

	text := handler.Clean(inputStr)

	return confirmBeforePost(roomInfo, text)

}

func postDraftHandler() error {
	roomInfo, err := getCurrentRoom()
	if err != nil {
		return err
	}
	draftFile := draft.DraftFile{}
	draftFile.Read()
	inputFromDraft := draftFile.Body
	text := handler.Clean(inputFromDraft)
	return confirmBeforePost(roomInfo, text)
}

// func openDraftHandler(inputStr string) error {
// 	// ドラフトファイルを開くだけ
// 	vim := vim.Vim{}
// 	vim.OpenDraftFile()
// 	return nil
// }
