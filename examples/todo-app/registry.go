package main

import "fmt"

type CommandRegistry struct {
	Commands          map[int]Command
	ValidCommandCodes []int
}

func (cr CommandRegistry) IsValidCommand(input int) bool {
	for _, i := range cr.ValidCommandCodes {
		if input == i {
			return true
		}
	}
	return false
}

func (cr *CommandRegistry) RegisterCommand(code int, cmd Command) error {
	if cr.IsValidCommand(code) {
		return fmt.Errorf("Command with code %v already registered", code)
	}
	cr.ValidCommandCodes = append(cr.ValidCommandCodes, code)
	cr.Commands[code] = cmd
	return nil
}

func (cr CommandRegistry) GetCommand(code int) (Command, error) {
	command, in := cr.Commands[code]
	if !in {
		return Command{}, fmt.Errorf("%v is not a valid command", code)
	}
	return command, nil
}

func (cr CommandRegistry) PrintMenu() {
	fmt.Printf("Menu:\n")
	for _, commandCode := range cr.ValidCommandCodes {
		command, _ := cr.GetCommand(commandCode)
		fmt.Printf("%v. %v\n", commandCode, command.Name)
	}
}
