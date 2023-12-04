package website

import (
	"plezk/lib/common"
	"plezk/lib/domain"
)

func Edit(dom domain.Domain) int {
	sslEnabled := dom.Info().SSL
	sshEnabled := dom.Info().HasSshAccess()
	menu_items := [][]string{}
	if sslEnabled {
		menu_items = append(menu_items, []string{"Disable SSL", "disable_ssl"})
	} else {
		menu_items = append(menu_items, []string{"Enable SSL", "enable_ssl"})
	}
	if sshEnabled {
		menu_items = append(menu_items, []string{"Disable SSH", "disable_ssh"})
	} else {
		menu_items = append(menu_items, []string{"Enable SSH", "enable_ssh"})
	}

	menu_items = append(menu_items, []string{"Back", "back"})
	choice := common.RenderBubbleTeaMenu(menu_items, "Edit "+dom.Name)

	back := 0

	switch choice {
	case "disable_ssl":
		dom.SSLDisable()
		return Edit(dom)
	case "enable_ssl":
		dom.SSLEnable()
		return Edit(dom)
	case "disable_ssh":
		dom.SshDisable()
		return Edit(dom)
	case "enable_ssh":
		dom.SshEnable()
		return Edit(dom)
	case "back":
		return back
	}

	return -1
}
