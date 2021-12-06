package main

import "strings"

type Command struct {
	Name string
	Args []string
	Len  int
	Full string
}

func CommandFromString(all string) *Command {
	split := strings.Split(all, " ")
	cmd := strings.ToLower(strings.ReplaceAll(split[0], "/", ""))
	var args []string
	if len(split) > 1 {
		args = split[1:]
	}
	return &Command{cmd, args, len(args), all}
}

func (cmd *Command) GetName() string {
	return cmd.Name
}

func (cmd *Command) GetArgs() []string {
	return cmd.Args
}

func (cmd *Command) GetLen() int {
	return cmd.Len
}

func (cmd *Command) FullString() string {
	return cmd.Full
}

func (cmd *Command) IssetArg(arg int) bool {
	return len(cmd.GetArgs()) >= (arg + 1)
}
