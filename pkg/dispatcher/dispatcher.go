package dispatcher

import (
	"errors"
	"fmt"
	"practice-telegram-bot/pkg/types"
)

type Dispatcher struct {
	Commands map[string]Command
	Chains   []Chain

	ChainCommandMap  map[string]Chain
	CurrentChain     map[int]Chain
	CurrentChainLink map[int]int
}

type Command struct {
	Name    string
	Handler func(update types.Update)
}

type Chain struct {
	StartChainCommand Command
	Handlers          []func(event types.Update)
}

func NewDispatcher() Dispatcher {
	return Dispatcher{Commands: make(map[string]Command), Chains: make([]Chain, 0), ChainCommandMap: make(map[string]Chain), CurrentChainLink: make(map[int]int), CurrentChain: make(map[int]Chain)}
}

func (d *Dispatcher) AddChain(chain Chain) {
	d.Chains = append(d.Chains, chain)
	d.AddCommand(chain.StartChainCommand)

	d.ChainCommandMap[chain.StartChainCommand.Name] = chain
}

func (d *Dispatcher) AddChains(chains []Chain) {
	for _, chain := range chains {
		d.AddChain(chain)
	}
}

func (d *Dispatcher) StartChain(command Command, user types.User) {
	d.CurrentChain[user.ID] = d.ChainCommandMap[command.Name]
	d.CurrentChainLink[user.ID] = 1
}

func (d *Dispatcher) NextChainStep(user types.User) error {
	d.CurrentChainLink[user.ID] += 1

	if d.CurrentChainLink[user.ID] > len(d.CurrentChain[user.ID].Handlers) {
		return errors.New("next step unreachable")
	}

	return nil
}

func (d *Dispatcher) ClearCurrentChain(user types.User) {
	d.CurrentChain[user.ID] = Chain{}
	d.CurrentChainLink[user.ID] = 0
}

func (d *Dispatcher) AddCommand(command Command) {
	d.Commands[fmt.Sprintf("/%s", command.Name)] = command
}

func (d *Dispatcher) AddCommands(commands []Command) {
	for _, command := range commands {
		d.AddCommand(command)
	}
}

func (d *Dispatcher) GetCommandByName(name string) (Command, error) {
	command, ok := d.Commands[name]

	if ok {
		return command, nil
	}

	return Command{}, errors.New("Command not found")
}
