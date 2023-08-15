package cli

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Print(`This program runs RGB backend server.
	 Usage:
	 
	 rgb [arguments]
	 
	 Supported arguments:

	`)
	flag.PrintDefaults()
	os.Exit(1)
}

// We set usage text that will be printed if app is started with some invalid options or arguments, after that we set default value and description for -env option.
// Finally, parse all CLI flags with flag.Parse().
// For the moment we don’t do anything smart with that value, we are only printing it.
// But that will change latter when we add logging.
func Parse() {
	flag.Usage = usage
	env := flag.String("env", "dev", `Sets run environment. Possible values are "dev" and "prod"`)
	flag.Parse()
	fmt.Println(*env)
}
