package main

import (
	"flag"
	"github.com/lepra-tsr/gdbt/util"
	"github.com/lepra-tsr/gdbt/validator"
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

	// parse command
	command := flag.Arg(0)
	err := validator.ValidateCommand(command)
	if err != nil {
		util.Echo("error!")
		util.Report(err)
	}
}
