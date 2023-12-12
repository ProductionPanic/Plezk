package websites

import (
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	"plezk/lib/colors"
	"plezk/lib/domain"
)

type DomainsAndWebsitesModel struct {
	Width          int
	Height         int
	domains        []domain.Domain
	cursor         int
	selectedDomain int
}

func (m DomainsAndWebsitesModel) Init() tea.Cmd {
	m.domains = domain.List()
	return nil
}

func (m *DomainsAndWebsitesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if m.cursor < len(m.domains)-1 {
				m.cursor++
			} else {
				m.cursor = len(m.domains) - 1
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = 0
			}
		case "enter":
			m.selectedDomain = m.cursor
		}
	}

	return m, nil
}

func (m DomainsAndWebsitesModel) View() string {
	rootStyle := lg.NewStyle().Width(m.Width).Height(m.Height).Background(lg.Color("#000000"))
	domainStyle := lg.NewStyle().Width(m.Width).Height(1).Padding(0, 0, 0, 0).Background(lg.Color("#000000"))
	cursorStyle := lg.NewStyle().Width(m.Width).Height(1).Padding(0, 0, 0, 0).Background(lg.Color("#ff3399"))
	cursor := cursorStyle.Render("")
	domains := ""
	for i, dom := range m.domains {
		if i == m.cursor {
			domains += cursor
		}
		domains += domainStyle.Render(dom.Name)
	}

	header := lg.NewStyle().BorderBackground(colors.Accent).Foreground(colors.AccentText).Width(m.Width-2).Height(1).Padding(0, 0, 0, 0).Border(lg.NormalBorder()).Render("Domains")

	return rootStyle.Render(header + domains)
}
