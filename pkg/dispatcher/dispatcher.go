package dispatcher

import (
	"errors"
	"fmt"
	"practice-telegram-bot/pkg/api"
	"practice-telegram-bot/pkg/types"
)

type Command struct {
	Name    string
	Handler func(event types.Update, api api.API)
}

type Command1 struct {
	Name    string
	Handler func(event types.Update, api api.API)
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
