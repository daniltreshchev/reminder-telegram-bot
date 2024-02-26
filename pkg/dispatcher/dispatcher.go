package dispatcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"practice-telegram-bot/pkg/state"
	"practice-telegram-bot/pkg/types"

	"github.com/redis/go-redis/v9"
)

type Dispatcher struct {
	Commands         map[string]Command `json:"commands"`
	Chains           map[string]Chain   `json:"chains"`
	ChainCommandMap  map[string]Chain   `json:"chain_command_map"`
	CurrentChain     map[int]string     `json:"current_chain"`
	CurrentChainLink map[int]int        `json:"current_chain_link"`
	useRedis         bool               `json:"-"`
	redisClient      *redis.Client      `json:"-"`
	backgroudContext context.Context    `json:"-"`
}

type Command struct {
	Name    string                    `json:"name"`
	Handler func(update types.Update) `json:"-"`
}

type Chain struct {
	Name              string                     `json:"name"`
	StartChainCommand Command                    `json:"start_chain_command"`
	Handlers          []func(event types.Update) `json:"-"`
}

func NewDispatcher() Dispatcher {
	return Dispatcher{
		Commands:         make(map[string]Command),
		Chains:           make(map[string]Chain),
		ChainCommandMap:  make(map[string]Chain),
		CurrentChainLink: make(map[int]int),
		CurrentChain:     make(map[int]string),
		useRedis:         false,
		backgroudContext: context.Background(),
	}
}

func NewRedisDispatcher(redisUrl string) Dispatcher {
	redisClient, err := state.NewRedisClient(redisUrl)

	if err != nil {
		fmt.Println("redis failds", err)
		return NewDispatcher()
	}

	redisDispatcher := NewDispatcher()
	redisDispatcher.redisClient = redisClient
	redisDispatcher.useRedis = true

	return redisDispatcher
}

func (d *Dispatcher) SaveRedisState() {
	if !d.useRedis {
		return
	}

	dispatcherBinary, err := json.Marshal(d)

	if err != nil {
		fmt.Println(err)
		return
	}

	d.redisClient.Set(d.backgroudContext, "dispatcher_state", dispatcherBinary, 0)
}

func (d *Dispatcher) GetRedisState() {
	if !d.useRedis {
		return
	}

	dispatcherBinary, err := d.redisClient.Get(d.backgroudContext, "dispatcher_state").Result()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal([]byte(dispatcherBinary), &d)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func (d *Dispatcher) AddChain(chain Chain) {
	d.GetRedisState()
	d.Chains[chain.Name] = chain
	d.AddCommand(chain.StartChainCommand)
	d.ChainCommandMap[chain.StartChainCommand.Name] = chain
}

func (d *Dispatcher) AddChains(chains []Chain) {
	for _, chain := range chains {
		d.AddChain(chain)
	}
}

func (d *Dispatcher) StartChain(command Command, user types.User) {
	d.GetRedisState()
	d.CurrentChain[user.ID] = d.ChainCommandMap[command.Name].Name
	d.CurrentChainLink[user.ID] = 0
	d.SaveRedisState()
}

func (d *Dispatcher) NextChainStep(user types.User) error {
	d.GetRedisState()
	d.CurrentChainLink[user.ID] += 1
	if d.CurrentChainLink[user.ID] > len(d.Chains[d.CurrentChain[user.ID]].Handlers) {
		return errors.New("next step unreachable")
	}
	d.SaveRedisState()
	return nil
}

func (d *Dispatcher) ClearCurrentChain(user types.User) {
	d.GetRedisState()
	delete(d.CurrentChain, user.ID)
	delete(d.CurrentChainLink, user.ID)
	d.SaveRedisState()
}

func (d *Dispatcher) AddCommand(command Command) {
	d.GetRedisState()
	d.Commands[fmt.Sprintf("/%s", command.Name)] = command
	d.SaveRedisState()
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
