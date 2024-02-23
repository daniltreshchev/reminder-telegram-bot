package services

import (
	"errors"
	"fmt"
	"practice-telegram-bot/pkg/botApi"
)

type Command struct {
	Name    string
	Handler func(event botApi.Update, api botApi.BotAPI)
}

type Dispatcher struct {
	Commands map[string]Command
}

func (d Dispatcher) AddCommand(command Command) {
	d.Commands[fmt.Sprintf("/%s", command.Name)] = command
}

func (d Dispatcher) GetCommandByName(name string) (Command, error) {
	command, ok := d.Commands[name]

	if ok {
		return command, nil
	}

	return Command{}, errors.New("Command not found")
}
