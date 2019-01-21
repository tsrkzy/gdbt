package post

import (
	"errors"
	"fmt"
)

const (
	EditorMode     = "editor"
	DirectPostMode = "directPost"
	OpenDraftMode  = "openDraft"
	PostDraftMode  = "postDraft"
)

func Handler(inputStr string, mode string) error {
	fmt.Println("post handler(" + mode + "): " + inputStr)

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

	// // オプションを取得
	// 	// 何もない場合 → vi起動し、可能なら引数を書き込んだ状態で表示する
	// 	// !# で始まる行は無視……説明を書く

	// 	// -m 標準入力 引数を使用する
	// 	// -d または --draft ドラフトファイルを使用

	// 	fpath := os.TempDir() + "/thetemporaryfile.txt"
	// 	f, err := os.Create(fpath)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	f.Close()

	// 	cmd := exec.Command("vi", fpath)
	// 	cmd.Stdin = os.Stdin
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Stderr = os.Stderr
	// 	err = cmd.Start()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	err = cmd.Wait()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println("Successfully edited.")
	// 	}

	return nil
}

func editorHandler(inputStr string) error {
	fmt.Println(" -> editorHandler.")
	return nil
}
func directPostHandler(inputStr string) error {
	fmt.Println(" -> directPostHandler.")
	return nil
}
func openDraftHandler(inputStr string) error {
	fmt.Println(" -> openDraftHandler.")
	return nil
}
func postDraftHandler(inputStr string) error {
	fmt.Println(" -> postDraftHandler.")
	return nil
}
