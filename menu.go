package main

import (
	"errors"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	"plezk/lib/common"
)

type Menu struct {
	Items    []MenuItem
	Cursor   int
	Selected int
	Focused  bool
}

type MenuItem struct {
	Name  string
	Type  int
	model string
}

func (m *Menu) Init() tea.Cmd {
	m.Cursor = 0
	m.Selected = -1
	m.Focused = true
	return nil
}

func (m *Menu) TotalLength() int {
	return len(m.Items)
}

func (m *Menu) GetSelected() *MenuItem {
	if m.Selected == -1 {
		return nil
	}
	return &m.Items[m.Selected]
}

func (m *Menu) GetModel() (string, error) {
	if m.Selected == -1 {
		return "", errors.New("No model selected")
	}
	return m.Items[m.Selected].model, nil
}

func (m *Menu) Up() {
	if m.Cursor > 0 {
		m.Cursor--
	}
}

func (m *Menu) Down() {
	if m.Cursor < m.TotalLength()-1 {
		m.Cursor++
	}
}

func (m *Menu) Unselect() {
	m.Selected = -1
	m.Focused = true
}

func (m *Menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			m.Up()
		case "down", "j":
			m.Down()
		case "enter", "right", "l":
			return m, common.SelectModel(m.Items[m.Cursor].model)
		}
	}
	return m, nil
}

func (m Menu) View() string {
	s := []string{}
	s = append(s, lg.NewStyle().Bold(true).Background(lg.Color("#000000")).Foreground(lg.Color("#FF00FF")).Render("Plezk"))
	for i, item := range m.Items {
		if i == m.Cursor {
			s = append(s, "> ")
		} else {
			s = append(s, "  ")
		}
		s[len(s)-1] += item.Name
	}
	return lg.JoinVertical(lg.Left, s...)
}
