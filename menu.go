package main

import (
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	"plezk/lib/colors"
)

type PleskMenu struct {
	Items         []PleskMenuItem
	CursorIndex   int
	SelectedIndex int
	MenuFocused   bool
	width         int
	height        int
}

func (m PleskMenu) Init() tea.Cmd {
	return nil
}

func (m PleskMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.MenuFocused {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "up", "k":
				if m.CursorIndex > 0 {
					m.CursorIndex--
				}
			case "down", "j":
				if m.CursorIndex < len(m.Items)-1 {
					m.CursorIndex++
				}
			case "enter":
				m.SelectedIndex = m.CursorIndex
				m.MenuFocused = false
				return m, nil
			}
		}
	}

	return m, nil
}

func (m PleskMenu) View() string {
	s := ""
	menuItemStyle := lg.NewStyle().
		Foreground(colors.Primary).
		Background(colors.Black).
		Width(m.width-2).Padding(0, 1)
	activeMenuItemStyle := menuItemStyle.Copy().
		Foreground(colors.PrimaryText).
		Background(colors.Primary)
	s_ := lg.NewStyle().
		Foreground(colors.PrimaryText).
		Background(colors.Black).
		Width(m.width).
		Align(lg.Center).
		MarginTop(1).
		Bold(true)

	s += s_.Render("Plesk")
	s += s_.Bold(false).Faint(true).Render("A plesk helper cli")
	s += "\n\n"

	for i, item := range m.Items {
		if i > 0 {
			s += "\n"
		}
		style := menuItemStyle.Copy()
		if m.SelectedIndex == i {
			style = activeMenuItemStyle.Copy()
		}
		if i == m.CursorIndex && m.MenuFocused {
			style = style.Copy().Border(lg.InnerHalfBlockBorder()).BorderRight(false).BorderTop(false).BorderBottom(false).BorderForeground(lg.Color("#fff"))
		} else {
			style = style.Copy().Border(lg.InnerHalfBlockBorder()).BorderRight(false).BorderTop(false).BorderBottom(false).BorderForeground(lg.Color("#000"))
		}

		s += style.
			Render(item.Label)
	}

	return lg.NewStyle().
		Height(m.height).
		Background(colors.Black).
		Border(lg.BlockBorder()).
		BorderRight(true).BorderLeft(false).BorderTop(false).BorderBottom(false).
		BorderForeground(colors.Primary).
		Width(m.width).Render(
		lg.Place(m.width, m.height, lg.Center, lg.Top, s))
}

type PleskMenuItem struct {
	Label             string
	PleskMenuItemType int // MenuItemTypes.*
	model             tea.Model
	function          func()
	command           tea.Cmd
}
