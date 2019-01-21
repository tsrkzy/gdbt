package main

import (
	// "bufio"
	"errors"
	"fmt"
	"github.com/lepra-tsr/gdbt/handler/list"
	// "github.com/lepra-tsr/gdbt/handler/post"
	"github.com/lepra-tsr/gdbt/handler/room"
	"github.com/lepra-tsr/gdbt/handler/setup"
	"github.com/urfave/cli"
	"os"
	"os/exec"
)

func main() {
	app := cli.NewApp()

	app.Name = "gdbtApp"
	app.Usage = "idobata unofficial cli tool"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {

		switch cmd := context.Args().Get(0); cmd {
		case "i":
			fallthrough
		case "init":
			if err := setup.Handler(); err != nil {
				fmt.Println(err)
			}

		case "r":
			fallthrough
		case "room":
			if err := room.Handler(); err != nil {
				fmt.Println(err)
			}

		case "l":
			fallthrough
		case "list":
			if err := list.Handler(); err != nil {
				fmt.Println(err)
			}

		case "p":
			fallthrough
		case "post":
			// オプションを取得
			// 何もない場合 → vi起動し、可能なら引数を書き込んだ状態で表示する
			// !# で始まる行は無視……説明を書く
			
			// -m 標準入力 引数を使用する
			// -d または --draft ドラフトファイルを使用

			fpath := os.TempDir() + "/thetemporaryfile.txt"
			f, err := os.Create(fpath)
			if err != nil {
				fmt.Println(err)
			}
			f.Close()

			cmd := exec.Command("vi", fpath)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Start()
			if err != nil {
				fmt.Println(err)
			}
			err = cmd.Wait()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Successfully edited.")
			}

			// if err := post.Handler(); err != nil {
			// 	fmt.Println(err)
			// }

			// buf := bufio.NewReader(os.Stdin)
			// if sentence, err := buf.ReadBytes('\n'); err != nil {
			// 	fmt.Println(err)
			// } else {
			// 	fmt.Println(string(sentence))
			// 	if err := post.Handler(string(sentence)); err != nil {
			// 		fmt.Println(err)
			// 	}
			// }
		default:
			fmt.Println("invalid command: " + cmd)
			return errors.New("invalid command: " + cmd)
		}

		/* exit correctly */
		return nil
	}

	app.Run(os.Args)
}
