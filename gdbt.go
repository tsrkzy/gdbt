package main

import (
	"errors"
	"flag"
	""
)

func main() {
	/** usage
	 * use hello [-flags value [-otherFlags otherValue]] args 
	 */
	// var (
	// 	showHelp = flag.Bool("help", false, "show help")
	// )
	flag.Parse()
	// fmt.Println(*showHelp)

	command := flag.Arg(0)
	err := validateCommand(command)
	if err != nil {
		echo("error!")
		report(err)
	}
}





func validateCommand(commandStr string) error {
	commandList := []string{"init", "channel", "list"}
	i := strIndexOf(commandStr, commandList)
	if i == -1 {
		return errors.New("validation error: invalid command: \"" + commandStr + "\"")
	}

	return nil
}
