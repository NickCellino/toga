package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"toga/parse"

	"github.com/mitchellh/cli"
)

type EvalCommand struct{}

func (EvalCommand) Synopsis() string {
	return "Welcome to the eval command."
}

func (ec EvalCommand) Name() string {
	return "eval"
}

func (c EvalCommand) Run(args []string) int {
	var rule, context string

	flagSet := flag.NewFlagSet(c.Name(), flag.ContinueOnError)
	flagSet.StringVar(&rule, "rule", "", "")
	flagSet.StringVar(&context, "context", "", "")
	flagSet.Parse(args)

	fmt.Println("rule received by eval:", rule)
	var data interface{}
	json.Unmarshal([]byte(rule), &data)
	ruleExpression, err := parse.ConvertToAst(data)
	if err != nil {
		fmt.Printf("Received error: %v\n", err)
	}
	fmt.Println("parsed rule expression: ", ruleExpression)

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
