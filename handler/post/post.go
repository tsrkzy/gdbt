package post

import "fmt"

func Handler(sentence string) error {
	fmt.Println("post handler: " + sentence)
	return nil
}
