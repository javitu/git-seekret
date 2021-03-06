package main

import(
	"fmt"
	"github.com/urfave/cli"
	"github.com/libgit2/git2go"
)

func GitSeekretRules(c *cli.Context) error {
	err := gs.LoadConfig(true)
	if git.IsErrorClass(err, git.ErrClassConfig) {
		return fmt.Errorf("Config not initialised - Try: 'git-seekret config --init'")
	}
	if err != nil {
		return err
	}	

	enable := c.String("enable")
	disable := c.String("disable")

	if enable != "" {
		err := gs.EnableRule(enable)
		if err != nil {
			return err
		}
	}

	if disable != "" {
		err := gs.DisableRule(disable)
		if err != nil {
			return err
		}
	}

	fmt.Println("List of rules:")
	for _, r := range gs.seekret.ListRules() {
		status := " "
		if r.Enabled {
			status = "x"
		}
		fmt.Printf("\t[%s] %s\n", status, r.Name)
	}

	gs.SaveConfig()
	
	return nil
}