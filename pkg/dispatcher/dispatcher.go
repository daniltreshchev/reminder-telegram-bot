package dispatcher

import (
	"context"
	"errors"
	"fmt"
	"practice-telegram-bot/pkg/state"
	"practice-telegram-bot/pkg/types"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Dispatcher struct {
	Commands         map[string]Command
	Chains           map[string]Chain
	ChainCommandMap  map[string]Chain
	CurrentChain     map[int]string
	CurrentChainLink map[int]int

	useRedis         bool
	redisClient      *redis.Client
	backgroudContext context.Context
}

type Command struct {
	Name    string
	Handler func(update types.Update)
}

type Chain struct {
	Name              string
	StartChainCommand Command
	Handlers          []func(event types.Update)
}

func NewDispatcher() Dispatcher {
	return Dispatcher{
		Commands:         make(map[string]Command),
		Chains:           make(map[string][]Chain),
		ChainCommandMap:  make(map[string]Chain),
		CurrentChainLink: make(map[int]int),
		CurrentChain:     make(map[int]string),
		useRedis:         false,
	}
}

func NewRedisDispatcher(redisUrl string) Dispatcher {
	redisClient, err := state.NewRedisClient(redisUrl)

	if err != nil {
		fmt.Println("redis failds", err)
		return NewDispatcher()
	}

	return Dispatcher{
		Commands:         make(map[string]Command),
		Chains:           make(map[string][]Chain),
		ChainCommandMap:  make(map[string]Chain),
		CurrentChain:     make(map[int]string),
		CurrentChainLink: make(map[int]int),
		useRedis:         true,
		redisClient:      redisClient,
		backgroudContext: context.Background(),
	}
}

func (d *Dispatcher) AddChain(chain Chain) {
	d.Chains[chain.Name] = append(d.Chains[chain.Name], chain)
	d.AddCommand(chain.StartChainCommand)

	d.ChainCommandMap[chain.StartChainCommand.Name] = chain
}

func (d *Dispatcher) AddChains(chains []Chain) {
	for _, chain := range chains {
		d.AddChain(chain)
	}
}

func (d *Dispatcher) StartChain(command Command, user types.User) {
	d.CurrentChain[user.ID] = d.ChainCommandMap[command.Name].Name
	d.CurrentChainLink[user.ID] = 1
}

func (d *Dispatcher) NextChainStepRedis(user types.User) error {
	currentChainLinkOld, err := d.redisClient.Get(d.backgroudContext, strconv.Itoa(user.ID)).Result()

	if err != nil {
		return err
	}

	currentChainLink, _ := strconv.Atoi(currentChainLinkOld)

	// if currentChainLink+1 > len() {

	// }

	err = d.redisClient.Set(d.backgroudContext, strconv.Itoa(user.ID), currentChainLink+1, 0).Err()

	return err
}

func (d *Dispatcher) NextChainStep(user types.User) error {
	// err := NextChainStepRedis(user)

	d.CurrentChainLink[user.ID] += 1
	if d.CurrentChainLink[user.ID] > len(d.Chains[d.CurrentChain[user.ID]]) {
		return errors.New("next step unreachable")
	}

	return nil
}

func (d *Dispatcher) ClearCurrentChain(user types.User) {
	d.CurrentChain[user.ID] = ""
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
