package cmd

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
)

var Args = os.Args

// Command .
type Command struct {
	// set flags
	flags *flag.FlagSet
	// Use is the one-line usage message.
	Use string
	// Short is the short description shown in the 'help' output.
	Short string
	// Long is the long message shown in the 'help <this-command>' output.
	Long string
	// SetOptions:
	SetOptions func(c *Command) error
	// Parse:
	Parse func(c *Command) error
	// Run: Typically the actual work function. Most commands will only implement this.
	Run func(cmd *Command, args []string)
}

// Execute .
func (c *Command) Execute() error {
	if ok := c.SetOptions(c); ok != nil {
		fmt.Println("Error in SetOptions!")
		return ok
	}
	if ok := c.Parse(c); ok != nil {
		fmt.Println("Error in Parsing!")
		return ok
	}
	c.Run(c, Args)
	return nil
}

func (c *Command) Flags() *flag.FlagSet {
	if c.flags == nil {
		c.flags = flag.NewFlagSet(c.Use, flag.ContinueOnError)
	}
	return c.flags
}
