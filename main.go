package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"toga/eval"
	"toga/parse"

	"github.com/mitchellh/cli"
)

type EvalCommand struct{}

func (EvalCommand) Synopsis() string {
	return "evaluates a rule"
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

	logger := log.Default()
	logger.Printf("evaluating: %v", rule)
	var data interface{}
	json.Unmarshal([]byte(rule), &data)
	if data == nil {
		logger.Fatalf("invalid json: %v\n", rule)
	}
	ruleExpression, err := parse.ConvertToAst(data)
	if err != nil {
		logger.Fatalf("received error: %v\n", err)
	}
	logger.Printf("parsed rule expression: %v", ruleExpression)

	value, err := ruleExpression.Eval(eval.Context{})
	if err != nil {
		logger.Fatalf("received error: %v\n", err)
	}
	logger.Printf("output: %v\n", value)

	return 0
}

func (c EvalCommand) Help() string {
	return `Usage:
toga eval -rule=<rule>

  Evaluates the provided rule.
`
}

func main() {
	c := cli.NewCLI("toga", "0.0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"eval": func() (cli.Command, error) {
			return EvalCommand{}, nil
		},
	}
	exitStatus, _ := c.Run()
	os.Exit(exitStatus)
}
