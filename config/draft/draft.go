package draft

import (
	. "github.com/lepra-tsr/gdbt/config"
	"io/ioutil"
	"os"
)

type DraftFile struct {
	Body string
}

func (u *DraftFile) Write() error {
	file, err := os.OpenFile(DraftFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0774)
	if err != nil {
		return err
	} else {
		defer file.Close()
		inputBytes := []byte(u.Body)
		file.Write(inputBytes)

		return nil
	}
}

func (u *DraftFile) Read() error {
	if bytes, err := ioutil.ReadFile(DraftFilePath); err != nil {
		return err
	} else {
		u.Body = string(bytes)
		return nil
	}
}
