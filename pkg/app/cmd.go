package app

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

type Command struct {
	usage    string
	desc     string
	options  CliOptions
	commands []*Command
	runFunc  RunCommandFunc
}

type RunCommandFunc func(args[] string) error

func (c *Command)cobraCommand()*cobra.Command{
	cmd := &cobra.Command{
		Use: c.usage,
		Short: c.desc,
	}
	cmd.SetOutput(os.Stdout)
	cmd.Flags().SortFlags = false
	if len(c.commands)>0{
		for _,comand := range c.commands{
			cmd.AddCommand(comand.cobraCommand())
		}
	}

	if c.runFunc != nil{
		cmd.Run = c.runCommand
	}
	//fmt.Println("this is cobracommand func",cmd)
	return cmd
}
func (c *Command)runCommand(cmd *cobra.Command,args[]string){
	if c.runFunc != nil{
		if err := c.runFunc(args);err != nil{
			fmt.Printf("%v %v\n", color.RedString("Error:"), err)
			os.Exit(1)
		}
	}
}