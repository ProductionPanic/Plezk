package common

import (
	"fmt"
	"github.com/ProductionPanic/go-pretty"
	tea "github.com/charmbracelet/bubbletea"
)

type BubbleTeaMainMenu struct {
	menuItems []string
	cursor    int
	title     string
}

func (m *BubbleTeaMainMenu) Init() tea.Cmd {
	return nil
}

func (m *BubbleTeaMainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.menuItems) - 1
			}
		case "down", "j":
			if m.cursor < len(m.menuItems)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
		case "enter", " ":
			return m, tea.Quit
		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *BubbleTeaMainMenu) View() string {
	s := "  [blue,bold]" + m.title + ":[]\n"
	for i, item := range m.menuItems {
		cursor := " "
		if m.cursor == i {
			cursor = "[cyan]>[]"
		}
		s += fmt.Sprintf("%s %s\n", cursor, item)
	}

	return pretty.Parse(s)
}

func RenderBubbleTeaMenu(menuItems [][]string, title string) string {
	menuItemKeys := make([]string, len(menuItems))
	menuItemValues := make([]string, len(menuItems))
	for i, item := range menuItems {
		menuItemKeys[i] = item[0]
		menuItemValues[i] = item[1]
	}
	p := tea.NewProgram(&BubbleTeaMainMenu{
		menuItems: menuItemKeys,
		cursor:    0,
		title:     title,
	})
	m, e := p.Run()
	if e != nil {
		panic(e)
	}
	return menuItemValues[m.(*BubbleTeaMainMenu).cursor]
}
