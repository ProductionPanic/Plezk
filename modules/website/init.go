package website

import (
	"fmt"
	"os"
	"os/exec"
	"plezk/lib/common"
	"plezk/lib/domain"

	tea "github.com/charmbracelet/bubbletea"
)

func Start(dom domain.Domain) {
	// Load some text for our viewport

	common.Cls()
	p := tea.NewProgram(
		model{content: info(dom), title: dom.Name},
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}

func info(selected domain.Domain) string {
	cmd := "plesk bin domain --info " + selected.Name
	cmdOut, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
		os.Exit(1)
	}
	return string(cmdOut)
}
