package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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
	var verbose bool

	flagSet := flag.NewFlagSet(c.Name(), flag.ContinueOnError)
	flagSet.StringVar(&ruleStr, "rule", "", "")
	flagSet.StringVar(&contextStr, "context", "", "")
	flagSet.BoolVar(&verbose, "verbose", false, "")
	flagSet.Parse(args)

	logger := log.Default()
	if !verbose {
		logger.SetOutput(ioutil.Discard)
	}
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

	var context map[string]interface{}
	contextExpression := eval.Context{}
	if contextStr != "" {
		err := json.Unmarshal([]byte(contextStr), &context)
		if err != nil {
			logger.Fatalf("error unmarshalling context json: %v\n", err)
		}
		logger.Printf("with context: %v", contextStr)
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

	rawValue, err := value.AsBool()
	if err != nil {
		logger.Fatalf("error serializing result value: %v\n", err)
	}

	output, err := json.Marshal(rawValue)
	if err != nil {
		logger.Fatalf("error marshalling json output: %v\n", err)
	}
	fmt.Println(string(output))

	return 0
}

func (c EvalCommand) Help() string {
	return `Usage:
toga eval -rule=<rule-json> -context=<context-json>

  Evaluates the provided rule.

  -rule <rule-json>
  The JSON rule to evaluate

  -context <context-json>
  The JSON context to use to evaluate the rule (default '{}')

  -verbose
  Whether to output verbose logging to stderr
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
