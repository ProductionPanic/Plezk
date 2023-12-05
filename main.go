package main

import (
	"fmt"
	"plezk/lib/common"
	"plezk/modules/tools"
	"plezk/modules/websites"
)

func main() {
	RenderMainMenu()
}

func RenderMainMenu() int {
	common.Cls()
	selected := common.RenderBubbleTeaMenu([][]string{
		{"Websites & domains", "websites"},
		{"Tools & settings", "tools"},
		{"Exit", "exit"},
	}, "Plezk")
	common.Cls()

	if selected == "exit" {
		common.Cls()
		fmt.Println("Bye!")
		return -1
	} else if selected == "websites" {
		r := websites.Start()
		if r == 0 {
			common.Cls()
			return RenderMainMenu()
		}
		return -1
	} else if selected == "tools" {
		common.Cls()
		return tools.Start()
	}
	common.Cls()
	return -1
}
