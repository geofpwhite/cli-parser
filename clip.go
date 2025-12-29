package main

import (
	"fmt"
	"strings"
)

type ClipParseError struct {
	invalidArgs []string
}
type ClipRunError struct {
	invalidArgs []string
	err         error
}

func (cre *ClipRunError) Error() string {
	return fmt.Sprintf(
		"Error executing command for arguments: %v\nError:%v",
		strings.Join(cre.invalidArgs, " "),
		cre.err,
	)
}

func (cpe *ClipParseError) Error() string {
	return fmt.Sprintf("Couldn't parse arguments: %v", strings.Join(cpe.invalidArgs, " "))
}

type Parser struct {
	cmds CommandNode
}

type CommandNode struct {
	subCommands map[string]*CommandNode
	cmd         func(...string) error
}

func (p *Parser) Parse(args []string) (func(...string) error, []string, bool) {
	cur := &p.cmds
	argIndex := 0
	for cur.subCommands[args[argIndex]] != nil {
		cur = cur.subCommands[args[argIndex]]
		argIndex++
	}
	return cur.cmd, args[argIndex:], cur.cmd != nil
}

func (p *Parser) Run(args []string) error {
	fn, args, ok := p.Parse(args)
	if !ok {
		return &ClipParseError{invalidArgs: args}
	}
	err := fn(args...)
	if err != nil {
		return &ClipRunError{err: err, invalidArgs: args}
	}
	return nil
}

func (p *Parser) AddCommand(args []string, cmd func(...string) error) {
	cur := &p.cmds
	for _, arg := range args {
		if cur.subCommands == nil {
			cur.subCommands = make(map[string]*CommandNode)
		}
		if cur.subCommands[arg] == nil {
			cur.subCommands[arg] = &CommandNode{}
		}
		cur = cur.subCommands[arg]
	}
	cur.cmd = cmd
}
