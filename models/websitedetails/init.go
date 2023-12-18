package websitedetails

import (
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	"plezk/lib/colors"
	"plezk/lib/common"
	"plezk/lib/domain"
)

type WebsiteDetailsModel struct {
	Domain *domain.Domain
	Width  int
	Height int
}

func (m *WebsiteDetailsModel) Init() tea.Cmd {
	return nil
}

func (m *WebsiteDetailsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "left", "h":
			return m, common.SelectModel("websites")
		}
	}
	return m, nil
}

func (m *WebsiteDetailsModel) View() string {
	info := lg.NewStyle().Width(m.Width/2-2).Background(colors.Black).Height(m.Height).Border(lg.BlockBorder(), false, false, false, true).BorderForeground(colors.Primary).BorderBackground(colors.Black).Render(
		lg.Place(m.Width/2-2, m.Height-2, lg.Center, lg.Center, lg.NewStyle().Align(lg.Left).Render(m.Domain.Info().GetInfoString())),
	)
	var actions []string
	// title
	title := lg.NewStyle().Width(m.Width).Align(lg.Center).Bold(true).Render(m.Domain.Name)
	actions = append(actions, title)
	// remove button
	removeButton := lg.NewStyle().Width(m.Width).Align(lg.Center).Render("Remove")
	actions = append(actions, removeButton)
	// back button
	backButton := lg.NewStyle().Width(m.Width).Align(lg.Center).Render("Back")
	actions = append(actions, backButton)

	actionsStr := lg.NewStyle().Foreground(colors.White).Background(colors.Black).Width(m.Width/2 - 2).Render(lg.JoinVertical(lg.Left, actions...))
	return lg.NewStyle().Width(m.Width).Height(m.Height).Render(lg.JoinHorizontal(lg.Left, actionsStr, info))
}
