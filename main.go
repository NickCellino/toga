package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

type EvalCommand struct{}

func (EvalCommand) Synopsis() string {
	return "Welcome to the eval command."
}

func (EvalCommand) Run(args []string) int {
	fmt.Printf("You have run the eval command with input: %v", args)
	return 0
}

func (EvalCommand) Help() string {
	return "sorry, no help here hombre."
}

func main() {
	c := cli.NewCLI("toga", "0.0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"eval": func() (cli.Command, error) {
			return EvalCommand{}, nil
		},
	}
	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println("There was an error!!", err)
	}
	os.Exit(exitStatus)
}
