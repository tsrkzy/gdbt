package vim

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type Vim struct {
}

func (u *Vim) OpenTemporaryFile(input string) (string, error) {
	tempFilePath, _ := filepath.Abs(filepath.Join(os.TempDir(), "/gdbt.md"))
	f, err := os.Create(tempFilePath)
	if err != nil {
		fmt.Println("failed to create temporary file: " + tempFilePath)
		return "", err
	}
	inputBytes := []byte(input)
	f.Write(inputBytes)
	f.Close()

	cmd := exec.Command("vim", tempFilePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		fmt.Println("cannot open temporary file with vim.")
		return "", err
	}
	err = cmd.Wait()
	if err != nil {
		return "", nil
	} else {
		bytes, err := ioutil.ReadFile(tempFilePath)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}
}
