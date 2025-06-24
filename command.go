package main

import ("errors"
				)

type Command struct {
	name string
	args []string
}


type Commands struct {
 commandsMap  map[string]func(*State, Command)(error)
}

func (c Commands) run(st *State, cmd Command) (error) {
	commandHandler, exists := c.commandsMap[cmd.name]
	if !exists {
		return errors.New("Command not found")
	}

	err := commandHandler(st, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c Commands) register(name string, f func(*State, Command)(error)) {
	c.commandsMap[name] = f
}
