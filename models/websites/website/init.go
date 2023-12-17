package website

import (
	tea "github.com/charmbracelet/bubbletea"
	"plezk/lib/domain"
)

type WebsiteModel struct {
	Domain domain.Domain
}

func (m *WebsiteModel) Init() tea.Cmd {
	return nil
}

type BackMsg struct{}

var Back = tea.Cmd(func() tea.Msg {
	return BackMsg{}
})

func (m *WebsiteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "left", "esc":
			return m, Back
		}
	}
	return m, nil
}

func (m *WebsiteModel) View() string {
	return m.Domain.Name
}
