package cmd

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cmd struct {
	Name string     // 命令名字
	Init func() int // 初始方法
	Exec func() int // 执行方法
	Desc string     // 描述信息
}

type Command struct {
	Name string
	cmds []*Cmd
}

func New(name string) *Command {
	return &Command{
		Name: name,
	}
}

func (c *Command) Get(name string) *Cmd {

	// 命令格式: name <command> [-options]
	name = strings.ToLower(name)
	for _, ci := range c.cmds {
		if strings.ToLower(ci.Name) == name {
			return ci
		}
	}

	return nil
}

func (c *Command) Add(name string, finit func() int, fexec func() int, desc string) *Command {
	c.cmds = append(c.cmds, &Cmd{
		Name: name,
		Init: finit,
		Exec: fexec,
		Desc: desc,
	})
	return c
}

func (c *Command) PrintUsage(w *os.File) {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "Useage:\n  %s <command> [-arguments]\nThe commands are:\n", c.Name)
	maxlen := 0
	for _, ci := range c.cmds {
		if ln := len(ci.Name); ln > maxlen {
			maxlen = ln
		}
	}
	for _, ci := range c.cmds {
		fmt.Fprintf(buf, "  %-"+strconv.Itoa(maxlen)+"s    %s\n", ci.Name, ci.Desc)
	}
	fmt.Fprintf(buf, "Use \"%s <command> -help\" for more information about a command.\n", c.Name)
	w.WriteString(buf.String())
}

var command *Command = New("")

func Exec(name string) int {
	// 命令格式: dataflow <command> [-arguments]
	if len(os.Args) > 1 {
		command.Name = name
		if c := command.Get(os.Args[1]); c != nil {
			phelp := flag.Bool("help", false, "print help information")
			if state := c.Init(); state != 0 {
				return state
			}
			// 重新解析一次! 否则help等参数无法解析到!
			if err := flag.CommandLine.Parse(os.Args[2:]); err != nil {
				fmt.Errorf("parse command line error: %v", err)
				return -1
			}

			if *phelp {
				fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", c.Name)
				flag.PrintDefaults()
				return 0
			}
			return c.Exec()
		}
	}
	// 打印帮助
	command.PrintUsage(os.Stdout)
	return 0
}

func Add(name string, finit func() int, fexec func() int, desc string) {
	command.Add(name, finit, fexec, desc)
}
