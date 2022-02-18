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
	var ruleStr, contextStr string

	flagSet := flag.NewFlagSet(c.Name(), flag.ContinueOnError)
	flagSet.StringVar(&ruleStr, "rule", "", "")
	flagSet.StringVar(&contextStr, "context", "", "")
	flagSet.Parse(args)

	logger := log.Default()
	logger.Printf("evaluating: %v", ruleStr)
	var rule interface{}
	err := json.Unmarshal([]byte(ruleStr), &rule)
	if err != nil {
		logger.Fatalf("invalid rule json: %v. error: %v\n", ruleStr, err)
	}
	ruleExpression, err := parse.ConvertToAst(rule)
	if err != nil {
		logger.Fatalf("received error: %v\n", err)
	}
	logger.Printf("parsed rule expression: %v", ruleExpression)

	var context map[string]interface{}
	contextExpression := eval.Context{}
	if contextStr != "" {
		err := json.Unmarshal([]byte(contextStr), &context)
		if err != nil {
			logger.Fatalf("error unmarshalling context json: %v\n", err)
		}
		logger.Printf("parsed context: %v", context)
		contextExpression, err = parse.ParseContext(context)
		if err != nil {
			logger.Fatalf("error parsing context: %v\n", err)
		}
	}

	value, err := ruleExpression.Eval(contextExpression)
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
