package draft

import (
	"errors"
	"fmt"

	confirmPrompt "github.com/lepra-tsr/gdbt/prompt/confirm"

	"github.com/lepra-tsr/gdbt/config/draft"
	"github.com/lepra-tsr/gdbt/config/room"
	"github.com/lepra-tsr/gdbt/handler"
	"github.com/lepra-tsr/gdbt/vim"
)

const (
	EditorMode    = "editor"
	PostDraftMode = "postDraft"
	Show          = "show"
)

func Handler(showFlag bool, postFlag bool) error {
	// fmt.Println("post handler(" + mode + "): " + inputStr)

	mode := EditorMode
	if postFlag {
		mode = PostDraftMode
	} else if showFlag {
		mode = Show
	} else {
		mode = EditorMode
	}

	switch mode {
	case EditorMode:
		if err := editorHandler(); err != nil {
			return err
		}
	case PostDraftMode:
		if err := postDraftHandler(); err != nil {
			return err
		}
	case Show:
		if err := showHandler(); err != nil {
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

func editorHandler() error {
	// ドラフトファイルを開くだけ
	vim := vim.Vim{}
	vim.OpenDraftFile()
	return nil
}

func showHandler() error {
	draftFile := draft.DraftFile{}
	draftFile.Read()
	fmt.Println(draftFile.Body)
	return nil
}

func postDraftHandler() error {
	roomInfo, err := getCurrentRoom()
	if err != nil {
		return err
	}
	draftFile := draft.DraftFile{}
	draftFile.Read()
	inputFromDraft := draftFile.Body

	return confirmBeforePost(roomInfo, inputFromDraft)
}

func confirmBeforePost(roomInfo *room.RoomInfo, text string) error {
	currentRoomId := roomInfo.Id
	currentConnectedName := roomInfo.GetConnectedName()

	cleanedText := handler.Clean(text)

	fmt.Println("- - - - - - - ")
	fmt.Println(cleanedText)
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
		handler.PostToRoom(cleanedText, currentRoomId)
		fmt.Println("post done.")
		return nil
	case "e":
		return editorHandler()
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
