package user

// init は予約語なので避けた

import (
	"fmt"

	. "github.com/lepra-tsr/gdbt/api/user"
)

func Handler() error {
	fmt.Println("user handler.")
	userJson := UserJson{}
	if err := userJson.Fetch(); err != nil {
		return err
	}

	userJson.Show()

	return nil
}
