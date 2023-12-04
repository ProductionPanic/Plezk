package websites

import (
	"plezk/lib/common"
	"plezk/lib/domain"
	"plezk/modules/website"

	"github.com/ProductionPanic/go-pretty"
)

const (
	MENU_T_SPLIT = "├"
	MENU_T_END   = "└"
)

func Start() {
	domains := domain.List()
	domains_menu := generate_domains_menu(domains, 0)
	menu := [][]string{}
	end := [][]string{
		{"Add domain", "add"},
		{"Add subdomain", "add_subdomain"},
		{"Add domain alias", "add_domain_alias"},
	}
	for _, d := range domains_menu {
		menu = append(menu, d)
	}
	for _, d := range end {
		menu = append(menu, d)
	}
	selected := common.RenderBubbleTeaMenu(menu, "Websites & domains")
	if selected == "add" {
		add_domain()
	} else {
		r := website.Start(domain.Get(selected))
		if r == 0 {
			Start()
		}
	}
}

func generate_domains_menu(domains []domain.Domain, depth int) [][]string {
	menu := make([][]string, 0)
	for i, domain := range domains {
		is_last := i == len(domains)-1
		prefix := ""
		if depth > 0 {
			for i := 0; i < depth; i++ {
				prefix += " "
			}
			prefix += "[cyan,bold]"
			if is_last {
				prefix += MENU_T_END + " [reset]"
			} else {
				prefix += MENU_T_SPLIT + " [reset]"
			}
		}
		menu = append(menu, []string{prefix + "[magenta,bold]" + domain.Name + "[]", domain.Name})
		if len(domain.Children) > 0 {
			subdomains_menu := generate_domains_menu(domain.Children, depth+1)
			for _, subdomain := range subdomains_menu {
				menu = append(menu, []string{subdomain[0], subdomain[1]})
			}
		}
	}
	return menu
}

func add_domain() {
	pretty.Println("Add domain")
}
