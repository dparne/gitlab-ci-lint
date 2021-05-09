package main

import (
	"fmt"

	"github.com/thatisuday/commando"
)

func main() {

	commando.
		SetExecutableName("gitlab-lint").
		SetVersion("0.0.1").
		SetDescription("This tool lints the .gitlab-ci.yml file against the Gitlab's API")

	commando.
		Register(nil).
		AddArgument("file", "File name if it's not .gitlab-ci.yml", ".gitlab-ci.yml").
		AddFlag("access-token, at", "level of depth to travel", commando.String, nil).
		AddFlag("project, p", "display size of the each file", commando.String, nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("Printing options of the `root` command...\n\n")

			// print arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			// print flags
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

			// configure info command
	commando.
		Register("info").
		SetShortDescription("gitlab ci linter").
		SetDescription("This tool lints the .gitlab-ci.yml file against the Gitlab's API").
		AddArgument("file", "local directory path", "./").
		AddFlag("level,l", "level of depth to travel", commando.Int, nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("Printing options of the `info` command...\n\n")

			// print arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			// print flags
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

	commando.Parse(nil)
}