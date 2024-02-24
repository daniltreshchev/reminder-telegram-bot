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
	CurrentChain     Chain
	CurrentChainLink int
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
	return Dispatcher{Commands: make(map[string]Command), Chains: make([]Chain, 0), ChainCommandMap: make(map[string]Chain), CurrentChainLink: -1}
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

func (d *Dispatcher) StartChain(command Command) {
	d.CurrentChain = d.ChainCommandMap[command.Name]
	d.CurrentChainLink = 0
}

func (d *Dispatcher) NextChainStep() error {
	d.CurrentChainLink += 1

	if d.CurrentChainLink > len(d.CurrentChain.Handlers) {
		return errors.New("next step unreachable")
	}

	return nil
}

func (d *Dispatcher) ClearCurrentChain() {
	d.CurrentChain = Chain{}
	d.CurrentChainLink = -1
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
