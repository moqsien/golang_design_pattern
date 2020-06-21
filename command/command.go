package command

import (
	"fmt"
)

// 命令模式
type Command interface {
	Execute()
}

type ConsoleOutput struct {
	message string
}

func (c *ConsoleOutput) Execute() {
	fmt.Println(c.message)
}

func CreateCommand(s string) Command {
	fmt.Println("Creating command")
	return &ConsoleOutput{
		message: s,
	}
}

type CommandQueue struct {
	queue []Command
}

func (q *CommandQueue) AddCommand(c Command) {
	q.queue = append(q.queue, c)
	if len(q.queue) == 3 {
		for _, command := range q.queue {
			command.Execute()
		}
	}
	if len(q.queue) == 3 {
		q.queue = make([]Command, 3)
	}
}
