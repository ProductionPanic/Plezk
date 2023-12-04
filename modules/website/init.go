package website

import (
	"plezk/lib/common"
	"plezk/lib/domain"
)

func Start(dom domain.Domain) int {
	// Load some text for our viewport

	common.Cls()

	r := common.RenderBubbleTeaMenu([][]string{
		{"View domain", "view"},
		{"Edit domain", "edit"},
		{"Delete domain", "delete"},
		{"Back", "back"},
	}, dom.Name)
	rc := -1
	switch r {
	case "view":
		rc = View(dom)
	case "edit":
		rc = Edit(dom)
	case "delete":
		rc = Delete(dom)
	case "back":
		return 0
	}
	if rc == 0 {
		return Start(dom)
	}
	return -1
}
