package main

import (
	"fmt"
	"plezk/lib/common"
	"plezk/modules/tools"
	"plezk/modules/websites"
)

func main() {
	common.Cls()
	selected := RenderBubbleTeaMainMenu([][]string{
		{"Websites & domains", "websites"},
		{"Tools & settings", "tools"},
		{"Exit", "exit"},
	})
	common.Cls()

	if selected == "exit" {
		fmt.Println("Bye!")
	} else if selected == "websites" {
		websites.Start()
	} else if selected == "tools" {
		tools.Start()
	}
}
