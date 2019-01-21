package post

import "fmt"

func Handler(sentence string) error {
	fmt.Println("post handler: " + sentence)

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
