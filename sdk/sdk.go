package sdk

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/NickCellino/toga/parse"
)

func EvalRuleFile(path string, context map[string]interface{}, value *interface{}) error {
	ruleFileContents, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error opening rule file: %w", err)
	}
	ruleStr := string(ruleFileContents)

	var rule interface{}
	err = json.Unmarshal([]byte(ruleStr), &rule)

	if err != nil {
		return fmt.Errorf("invalid rule json: %v. error: %v", ruleStr, err)
	}
	ruleExpression, err := parse.ConvertToAst(rule)
	if err != nil {
		return fmt.Errorf("error parsing rule: %v", err)
	}

	contextExpression, err := parse.ParseContext(context)
	if err != nil {
		return fmt.Errorf("error parsing context: %v", err)
	}
	evaledRule, err := ruleExpression.Eval(contextExpression)
	if err != nil {
		return fmt.Errorf("rule evaluation error: %w", err)
	}

	expectedKind := reflect.TypeOf(*value).Kind()
	if expectedKind == reflect.Bool {
		*value, _ = evaledRule.AsBool()
	} else if expectedKind == reflect.Float64 {
		*value, _ = evaledRule.AsNumber()
	} else if expectedKind == reflect.String {
		*value, _ = evaledRule.AsString()
	} else {
		return fmt.Errorf("value passed in should be either *bool, *float64 or *string")
	}

	return nil
}
