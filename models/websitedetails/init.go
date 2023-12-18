package websitedetails

import (
	tea "github.com/charmbracelet/bubbletea"
	"plezk/lib/domain"
)

type WebsiteDetailsModel struct {
	Domain domain.Domain
}

func (m *WebsiteDetailsModel) Init() tea.Cmd {
	return nil
}

func (m *WebsiteDetailsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *WebsiteDetailsModel) View() string {
	return m.Domain.Name
}
