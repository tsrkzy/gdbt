package setup
// init は予約語なので避けた

import "fmt"

func Handler() error {
	fmt.Println("init handler.")
	return nil
}
