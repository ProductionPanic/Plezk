package main

import (
	"fmt"
	"plezk/lib/admin"
	"plezk/lib/domain"

	"github.com/ProductionPanic/go-cursor"
	"github.com/ProductionPanic/go-input"
	"github.com/ProductionPanic/go-pretty"
)

var lastTree map[string]interface{}

func cls() {
	cursor.ClearScreen()
	cursor.Top()
}

var commandTree map[string]interface{}

func main() {
	commandTree = map[string]interface{}{
		"ssl": map[string]interface{}{
			"create": func() {
				selectedDomain := selectDomain()
				cls()
				pretty.Println("[bold,cyan]Selected domain:[]" + selectedDomain)
				pretty.Printf("[bold,white]Using admin account %s[]", admin.Info().Email)
				domain.SetSsl(selectedDomain, admin.Info().Email)
			},
			"back": treeBack,
		},
		"subdomain": map[string]interface{}{
			"create": func() {
				selectedDomain := selectDomain()
				cls()
				subdomainName := input.GetText("[bold,blue]Enter subdomain name:[]")
				cls()
				pretty.Println("[bold,cyan]Creating subdomain:[]" + subdomainName + "." + selectedDomain)
				success := domain.CreateSubdomain(subdomainName, selectedDomain)
				if success {
					pretty.Println("[bold,green]Success![]")
				} else {
					pretty.Println("[bold,red]Failed![]")
				}
			},
			"delete": func() {
				commandTree["domain"].(map[string]interface{})["delete"].(func())()
			},
			"back": treeBack,
		},
		"domain": map[string]interface{}{
			"list": func() {
				for i, d := range domain.List() {
					if len(d) == 0 {
						continue
					}
					pretty.Println(fmt.Sprintf("%d: %s", i, d))
				}
			},
			"create": func() {
				domainName := input.GetText("[bold,blue]Enter domain name:[]")
				cls()
				pretty.Println("[bold,white]Creating domain:[] " + domainName)
				success := domain.Create(domainName)
				if success {
					pretty.Println("[bold,green]Success![]")
				} else {
					pretty.Println("[bold,red]Failed![]")
				}
			},
			"delete": func() {
				selectedDomain := selectDomain()
				if !confirm("[bold,yellow]Are you sure you want to delete " + selectedDomain + "?[]") {
					return
				}

				cls()
				pretty.Println("[bold,cyan]Deleting domain:[]" + selectedDomain)
				success := domain.Delete(selectedDomain)
				if success {
					pretty.Println("[bold,green]Success![]")
				} else {
					pretty.Println("[bold,red]Failed![]")
				}
			},
			"back": treeBack,
		},
	}
	cursor.ClearScreen()
	cursor.Top()
	pretty.Println("[bold,blue]Plesk Helper[]")
	pretty.Println("[bold,blue]Version:[] 0.0.1")
	runTree(commandTree)
}

func runTree(tree map[string]interface{}) {
	var keys []string
	for k, _ := range tree {
		keys = append(keys, k)
	}
	selectBox := input.NewSelect("[bold,blue]Select a command:[]")
	for _, key := range keys {
		selectBox.AddItem(key, key)
	}
	selectedCommand := selectBox.Run()

	result := tree[selectedCommand.(string)]
	cursor.ClearScreen()
	cursor.Top()
	switch result.(type) {
	case func():
		result.(func())()
	case map[string]interface{}:
		lastTree = tree
		runTree(result.(map[string]interface{}))
	}
}

func treeBack() {
	runTree(lastTree)
}

func selectDomain() string {
	cls()
	domains := domain.List()
	selectBox := input.NewSelect("[bold,blue]Select a domain:[]")
	for _, domain := range domains {
		selectBox.AddItem(domain, domain)
	}
	return selectBox.Run().(string)
}

func confirm(prompt string) bool {
	pretty.Println(prompt + " (y/n)")
	for {
		var input string
		fmt.Scanln(&input)
		if input == "y" {
			return true
		} else if input == "n" {
			return false
		}
	}
}
