package main

import (
	"errors"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
	"plezk/models/websites"
)

type PlezkModel struct {
	Menu   *Menu
	Models map[string]tea.Model
}

func (m *PlezkModel) Init() tea.Cmd {
	return nil
}

func (m *PlezkModel) TotalLength() int {
	return len(m.Menu.Items)
}

func (m *PlezkModel) GetSelected() *MenuItem {
	if m.Menu.Selected == -1 {
		return nil
	}
	return &m.Menu.Items[m.Menu.Selected]
}

func (m *PlezkModel) GetModel() (tea.Model, error) {
	if m.Menu.Selected == -1 {
		return nil, errors.New("No model selected")
	}
	if !m.HasModel() {
		return nil, errors.New("Model does not exist")
	}
	model := m.Menu.Items[m.Menu.Selected].model
	return m.Models[model], nil
}

func (m *PlezkModel) HasModel() bool {
	selected := m.Menu.Selected
	if selected == -1 {
		return false
	}

	_, ok := m.Models[m.Menu.Items[selected].model]
	return ok
}

func (m *PlezkModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	if !m.Menu.Focused && !m.HasModel() {
		m.Menu.Focused = true
		return m, nil
	}

	if m.Menu.Focused {
		menum, cmd := m.Menu.Update(msg)
		m.Menu = menum.(*Menu)
		return m, cmd
	} else if m.HasModel() {
		var cmd tea.Cmd
		model := m.Models[m.Menu.Items[m.Menu.Selected].model]
		model, cmd = model.Update(msg)
		m.Models[m.Menu.Items[m.Menu.Selected].model] = model
		return m, cmd
	}

	return m, nil
}

func (m *PlezkModel) View() string {
	w, h, _ := term.GetSize(0)
	w -= 4
	h -= 2
	menu := lg.NewStyle().
		Width(w/6).
		Height(h).
		Background(lg.Color("#000000")).
		Foreground(lg.Color("#ffffff")).
		Padding(1, 2).
		Border(lg.OuterHalfBlockBorder()).
		BorderForeground(lg.Color("#9933ff")).
		BorderBackground(lg.Color("#000000")).Render(m.Menu.View())
	content_str := ""
	if m.HasModel() {
		content_str_model, _ := m.GetModel()
		content_str = content_str_model.View()
	}
	content := lg.NewStyle().
		Width(w-w/6).
		Height(h).
		Background(lg.Color("#000000")).
		Foreground(lg.Color("#000000")).
		Padding(1, 2).
		Border(lg.OuterHalfBlockBorder()).
		BorderForeground(lg.Color("#9933ff")).
		BorderBackground(lg.Color("#000000")).Render(content_str)

	return lg.NewStyle().Render(lg.JoinHorizontal(lg.Top, menu, content))
}

func main() {
	w, h, _ := term.GetSize(0)
	p := tea.NewProgram(&PlezkModel{
		Menu: &Menu{
			Items: []MenuItem{
				MenuItem{
					Name:  "Websites",
					Type:  0,
					model: "websites",
				},
				MenuItem{
					Name:  "Settings",
					Type:  0,
					model: "settings",
				},
			},
			Focused: true,
		},
		Models: map[string]tea.Model{
			"websites": &websites.DomainsAndWebsitesModel{
				Width:  w,
				Height: h,
			},
		},
	}, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		panic(err)
	}
}
