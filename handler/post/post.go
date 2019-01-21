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
	// テンポラリファイルを作成してvimで開く。
	// テンプレートを挿入。
	// 入力があったら追加する。
	// ファイルを閉じたら、テンポラリファイルを読み出す。
	// 末尾の改行を取り、コメントを削除。(共通処理？)
	//   本文が空ならエラー。
	// 本文を表示し、confirm。
	// enter ならば送信。
	//   e ならば再編集
	//   q ならば破棄して終了。
	//   dまたはそれ以外ならばドラフトを上書きして終了。

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
