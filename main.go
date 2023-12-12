package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
	"plezk/enum/MenuItemTypes"
	"plezk/lib/colors"
	"plezk/models/websites"
)
import lg "github.com/charmbracelet/lipgloss"

type PleskModel struct {
	Menu          PleskMenu
	MenuFocused   bool
	width         int
	height        int
	selectedModel tea.Model
}

type PleskMenu struct {
	Items         []PleskMenuItem
	CursorIndex   int
	SelectedIndex int
}

type PleskMenuItem struct {
	Label             string
	PleskMenuItemType int // MenuItemTypes.*
	model             tea.Model
	function          func()
	command           tea.Cmd
}

func (m PleskModel) Init() tea.Cmd {
	for i, item := range m.Menu.Items {
		if item.PleskMenuItemType == MenuItemTypes.ModelType {
			m.Menu.Items[i].model.Init()
		}
	}
	return nil
}

func (m *PleskModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	if m.MenuFocused {
		return m.UpdateMenu(msg)
	}

	if m.Menu.Items[m.Menu.SelectedIndex].PleskMenuItemType == MenuItemTypes.ModelType {
		return m.UpdateModel(msg)
	}

	return m, nil
}

func (m *PleskModel) UpdateModel(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.MenuFocused = true
			return m, nil
		}
	}
	var cmd tea.Cmd
	m.selectedModel, cmd = m.selectedModel.Update(msg)
	return m, cmd
}

func (m *PleskModel) UpdateMenu(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.Menu.CursorIndex > 0 {
				m.Menu.CursorIndex--
			}
		case "down", "j":
			if m.Menu.CursorIndex < len(m.Menu.Items)-1 {
				m.Menu.CursorIndex++
			}
		case "enter":
			m.MenuFocused = false
			switch m.Menu.Items[m.Menu.CursorIndex].PleskMenuItemType {
			case MenuItemTypes.ModelType:
				m.Menu.SelectedIndex = m.Menu.CursorIndex
				m.selectedModel = m.Menu.Items[m.Menu.CursorIndex].model
				return m, nil
			case MenuItemTypes.FunctionType:
				m.Menu.Items[m.Menu.CursorIndex].function()
			case MenuItemTypes.CommandType:
				return m, m.Menu.Items[m.Menu.CursorIndex].command
			}
		}
	}
	return m, nil
}

func (m *PleskModel) ViewMenu() string {
	s := lg.NewStyle().
		Foreground(colors.PrimaryText).
		Background(colors.Primary).
		Align(lg.Center).
		Padding(0, 2).
		Width(m.width / 6).
		Render("Plesk")
	menuItemStyle := lg.NewStyle().
		Foreground(colors.Primary).
		Background(colors.Black).
		Width(m.width/6-1).Padding(0, 1)
	activeMenuItemStyle := menuItemStyle.Copy().
		Foreground(colors.PrimaryText).
		Background(colors.Primary)
	s += "\n\n"
	for i, item := range m.Menu.Items {
		if i > 0 {
			s += "\n"
		}
		style := menuItemStyle.Copy()
		if m.Menu.SelectedIndex == i {
			style = activeMenuItemStyle.Copy()
		}
		if i == m.Menu.CursorIndex && m.MenuFocused {
			style = style.Copy().Border(lg.InnerHalfBlockBorder()).BorderRight(false).BorderTop(false).BorderBottom(false).BorderForeground(lg.Color("#fff"))
		} else {
			style = style.Copy().Border(lg.InnerHalfBlockBorder()).BorderRight(false).BorderTop(false).BorderBottom(false).BorderForeground(lg.Color("#000"))
		}

		s += style.Render(item.Label)
	}

	return lg.NewStyle().
		Height(m.height).
		Background(colors.Black).
		Width(m.width / 6).Render(
		lg.Place(m.width/6, m.height, lg.Center, lg.Top, s))
}

func (m *PleskModel) ViewContent() string {
	ContentStyle := lg.NewStyle().
		Foreground(lg.Color("205")).
		Background(colors.Black).
		Height(m.height - 1).
		Width(m.width / 6 * 5).
		Border(lg.NormalBorder()).
		BorderBackground(colors.Black).BorderForeground(colors.Primary).
		BorderRight(false).BorderTop(false).BorderBottom(false).
		Padding(1)
	pre := lg.NewStyle().Height(1).Width(m.width/6*5).Background(colors.Primary).Render("") + "\n"
	if m.Menu.SelectedIndex > 0 {
		pre += m.selectedModel.View()
	}
	return pre + ContentStyle.Render("")
}

func (m *PleskModel) View() string {
	w, h, _ := term.GetSize(0)
	menu := m.ViewMenu()
	content := m.ViewContent()
	return lg.Place(w, h, lg.Center, lg.Center, lg.JoinHorizontal(
		lg.Top,
		menu,
		content,
	))
}

type PlaceHolderModel struct {
	title  string
	width  int
	height int
}

func (m PlaceHolderModel) Init() tea.Cmd {
	return nil
}
func (m *PlaceHolderModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
func (m *PlaceHolderModel) View() string {
	return lg.NewStyle().Width(m.width).Height(m.height).Render(m.title)
}

func main() {
	program := tea.NewProgram(initialModel(), tea.WithAltScreen())
	_, err := program.Run()
	if err != nil {
		panic(err)
	}
}

func initialModel() tea.Model {
	w, h, _ := term.GetSize(0)
	w -= 10
	h -= 5
	return &PleskModel{
		Menu: PleskMenu{
			Items: []PleskMenuItem{
				{
					Label:             "Websites & domains",
					PleskMenuItemType: MenuItemTypes.ModelType,
					model: &websites.DomainsAndWebsitesModel{
						Width:  w/6*5 - 2,
						Height: h - 5,
					},
				},
				{
					Label:             "Files",
					PleskMenuItemType: MenuItemTypes.ModelType,
					model:             &PlaceHolderModel{width: w/6*5 - 2, height: h - 2, title: "Files"},
				},
				{
					Label:             "Databases",
					PleskMenuItemType: MenuItemTypes.ModelType,
					model:             &PlaceHolderModel{width: w/6*5 - 2, height: h - 2, title: "Databases"},
				},
				{
					Label:             "Mail",
					PleskMenuItemType: MenuItemTypes.ModelType,
					model:             &PlaceHolderModel{width: w/6*5 - 2, height: h - 2, title: "Mail"},
				},
				{
					Label:             "Tools & Settings",
					PleskMenuItemType: MenuItemTypes.ModelType,
					model:             &PlaceHolderModel{width: w/6*5 - 2, height: h - 2, title: "Tools & Settings"},
				},
				{
					Label:             "Docker",
					PleskMenuItemType: MenuItemTypes.ModelType,
					model:             &PlaceHolderModel{width: w/6*5 - 2, height: h - 2, title: "Docker"},
				},
				{
					Label:             "Exit",
					PleskMenuItemType: MenuItemTypes.CommandType,
					command:           tea.Quit,
				},
			},
		},
		MenuFocused: true,
		width:       w,
		height:      h,
	}
}
