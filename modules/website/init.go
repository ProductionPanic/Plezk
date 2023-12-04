package website

import (
	"plezk/lib/common"
	"plezk/lib/domain"
)

func Start(dom domain.Domain) {
	// Load some text for our viewport

	common.Cls()

	r := common.RenderBubbleTeaMenu([][]string{
		{"View domain", "view"},
		{"Edit domain", "edit"},
		{"Delete domain", "delete"},
		{"Back to main menu", "back"},
	}, dom.Name)

	switch r {
	case "view":
		View(dom)
	case "edit":
		Edit(dom)
	case "delete":
		Delete(dom)
	case "back":
		return
	}
}
