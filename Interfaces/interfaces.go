package main

import "fmt"

type Helper = interface {
	Help() string
}

type HelpString string

func (hs HelpString) Help() string {
	return string(hs)
}

type UnHelpString struct{}

func (uhs UnHelpString) Help() string {
	return "I can not help you"
}

var _ = Helper(HelpString(""))

func main() {
	var h Helper = HelpString("Help me")
	fmt.Println(h.Help())
	var explicit = interface{ Help() string }.Help(h)
	fmt.Println(explicit)

	var helpers = []Helper{
		HelpString("Help me again"),
		&UnHelpString{},
	}
	fmt.Println(helpers)
	for _, helper := range helpers {
		fmt.Println(helper.Help())
	}
}
