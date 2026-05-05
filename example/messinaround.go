package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/geofpwhite/clip"
)

func main() {
	c := config{ary: []int{1, 2, 3, 5, 4, 56, 7, 4, 2}}
	c.p.AddCommand([]string{"array"}, c.ArrayHelpText)
	c.p.AddSubCommand([]string{"array"}, "sum", c.SumCommand)
	err := c.p.Run(os.Args[1:])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

type config struct {
	ary []int
	p   clip.Parser
}

func (c *config) ArrayHelpText(args ...string) error {
	fmt.Printf("Array: %v", c.ary)
	return nil
}

func (c *config) SumCommand(args ...string) error {
	intArgs := make([]int, len(args))
	for i, numString := range args {
		n, err := strconv.Atoi(numString)
		if err != nil {
			return errors.New(args[0] + "can't be parsed as an int")
		}
		intArgs[i] = n
	}
	c.Sum()
	return nil
}

func (c *config) Sum() {
	sum := 0
	for _, num := range c.ary {
		sum += num
	}
	fmt.Println(sum)
}
