package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	c := config{}
	c.p.AddCommand([]string{"add"}, c.AddCommand)
	c.p.AddCommand([]string{"add", "negative"}, c.AddNegativeCommand)
	fn, args, ok := c.p.Parse([]string{"add", "23"})
	if !ok {
		panic("uh oh")
	}
	fn(args...)
	fn(args...)
	fmt.Println(c.x)
	err := c.p.Run(os.Args[1:])
	fmt.Println(err, c.x)
}

type config struct {
	x int
	p Parser
}

func (c *config) AddCommand(args ...string) error {
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New(args[0] + "can't be parsed as an int")
	}
	c.Add(n)
	return nil
}

func (c *config) AddNegativeCommand(args ...string) error {
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New(args[0] + "can't be parsed as an int")
	}
	c.Add(-n)
	return nil
}

func (c *config) Add(i int) {
	c.x += i
}
