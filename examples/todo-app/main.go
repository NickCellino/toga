package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/NickCellino/toga/sdk"
)

type TodoItem struct {
	Name string
	Done bool
}

type Command struct {
	Name   string
	Action func()
}

func PrintTodoItems(todoItems []TodoItem) {
	fmt.Printf("🗒️ Your todo list:\n")
	if len(todoItems) == 0 {
		fmt.Printf("No items yet!\n")
	} else {
		for idx, todoItem := range todoItems {
			var symbol string
			if todoItem.Done {
				symbol = "✅"
			} else {
				symbol = "❌"
			}
			fmt.Printf("%v: %v [%v]\n", idx, todoItem.Name, symbol)
		}
	}
	fmt.Println()
}

func GetUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSuffix(userInput, "\n")
	return userInput
}

func GetUserIntInput(prompt string) (int, error) {
	userInputStr := GetUserInput(prompt)
	userInput, err := strconv.Atoi(userInputStr)
	if err != nil {
		return 0, fmt.Errorf("expecting a number, got: %v", userInputStr)
	}
	return userInput, nil
}

func main() {
	fmt.Println("Welcome to Toga-Todo!")
	todoItems := []TodoItem{}
	commandRegistry := CommandRegistry{
		Commands: make(map[int]Command),
	}
	commandRegistry.RegisterCommand(1, Command{
		Name: "Add new todo items",
		Action: func() {
			newItem := GetUserInput("New item: ")
			todoItems = append(todoItems, TodoItem{Name: newItem, Done: false})
		},
	})
	commandRegistry.RegisterCommand(2, Command{
		Name: "Toggle todo item status",
		Action: func() {
			if len(todoItems) == 0 {
				fmt.Println("No items to toggle yet!")
				return
			}
			itemNumber, err := GetUserIntInput("Todo item to toggle: ")
			if err != nil {
				fmt.Printf("You must enter a number\n")
				return
			}
			if itemNumber >= len(todoItems) {
				fmt.Printf("You must enter a number below %v\n", len(todoItems))
				return
			}
			todoItem := &todoItems[itemNumber]
			todoItem.Done = !todoItem.Done
		},
	})
	regionVar := os.Getenv("REGION")
	context := map[string]interface{}{
		"region": regionVar,
	}
	removeItemEnabled := false
	err := sdk.EvalRuleFile("allow-remove-item.json", context, &removeItemEnabled)
	if err != nil {
		fmt.Printf("error calling toga sdk: %v\n", err)
	}
	if removeItemEnabled {
		commandRegistry.RegisterCommand(3, Command{
			Name: "Remove todo item",
			Action: func() {
				itemNumber, err := GetUserIntInput("Todo item to remove: ")
				if err != nil {
					fmt.Printf("You must enter a number\n")
					return
				}
				if itemNumber >= len(todoItems) {
					fmt.Printf("You must enter a number below %v\n", len(todoItems))
					return
				}
				todoItems = append(todoItems[:itemNumber], todoItems[itemNumber+1:]...)
			},
		})
	}
	for {
		fmt.Println()
		fmt.Print("----------------------------------------\n")
		PrintTodoItems(todoItems)
		fmt.Print("----------------------------------------\n")
		commandRegistry.PrintMenu()
		choice, err := GetUserIntInput(">> ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		command, err := commandRegistry.GetCommand(choice)
		if err != nil {
			fmt.Printf("error: %v", err)
			continue
		}
		command.Action()
		time.Sleep(300 * time.Millisecond)
	}
}
