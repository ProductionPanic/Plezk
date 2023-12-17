package main

import (
	"plezk/enum/MenuItemTypes"
	"plezk/lib/colors"
	"plezk/lib/common"
	"plezk/models/websites"

	tea "github.com/charmbracelet/bubbletea"

	lg "github.com/charmbracelet/lipgloss"
)

type PleskModel struct {
	Menu   PleskMenu
	width  int
	height int
}

func (m PleskModel) Init() tea.Cmd {
	for i, item := range m.Menu.Items {
		if item.PleskMenuItemType == MenuItemTypes.ModelType {
			m.Menu.Items[i].model.Init()
			m.Menu.Items[i].model.Update(nil)
		}
	}
	return nil
}
func (m *PleskModel) HasCurrentModel() bool {
	return m.Menu.Items[m.Menu.SelectedIndex].model != nil
}

func (m *PleskModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd

	if m.HasCurrentModel() && !m.Menu.MenuFocused {
		cur_model, cur_cmd := m.Menu.Items[m.Menu.SelectedIndex].model.Update(msg)
		m.Menu.Items[m.Menu.SelectedIndex].model = cur_model
		if cur_cmd != nil {
			cmds = append(cmds, cur_cmd)
		}
	}

	m_menu, m_cmd := m.Menu.Update(msg)
	m.Menu = m_menu.(PleskMenu)
	if m_cmd != nil {
		cmds = append(cmds, m_cmd)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
		}
	case common.BackMsg:
		m.Menu.MenuFocused = true
		return m, nil
	}
	return m, tea.Batch(cmds...)
}

func (m *PleskModel) UpdateModel(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.HasCurrentModel() && !m.Menu.MenuFocused {
		updatedModel, cmd := m.Menu.Items[m.Menu.SelectedIndex].model.Update(msg)
		m.Menu.Items[m.Menu.SelectedIndex].model = updatedModel
		return m, cmd
	}
	return m, nil
}

func (m *PleskModel) ViewContent() string {
	var modelOutput string
	if m.Menu.Items[m.Menu.SelectedIndex].PleskMenuItemType == MenuItemTypes.ModelType {
		modelOutput = m.Menu.Items[m.Menu.SelectedIndex].model.View()
	}
	return lg.NewStyle().Width(m.width / 6 * 5).Background(colors.Black).MaxHeight(m.height - 1).Height(m.height - 1).Render(modelOutput)
}

func (m *PleskModel) View() string {
	w, h := GetTermSize()
	menu := m.Menu.View()
	content := m.ViewContent()
	app := lg.NewStyle().
		Border(lg.InnerHalfBlockBorder()).
		BorderForeground(colors.Primary).
		Render(lg.JoinHorizontal(
			lg.Top,
			menu,
			content,
		))

	return lg.Place(w, h, lg.Center, lg.Center, app)
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
	return lg.NewStyle().Foreground(colors.Accent).Render(m.title)
}

func main() {
	program := tea.NewProgram(initialModel(), tea.WithAltScreen())
	_, err := program.Run()
	if err != nil {
		panic(err)
	}
}

func initialModel() tea.Model {
	w, h := GetTermSize()
	w -= 10
	h -= 5
	return &PleskModel{
		Menu: PleskMenu{
			SelectedIndex: 0,
			Items: []PleskMenuItem{
				{
					Label:             "Websites & domains",
					PleskMenuItemType: MenuItemTypes.ModelType,
					model: &websites.DomainsAndWebsitesModel{
						Width:          w / 6 * 5,
						Height:         h,
						SelectedDomain: -1,
					},
				},
				{
					Label:             "Files",
					PleskMenuItemType: MenuItemTypes.ModelType,
					model:             &PlaceHolderModel{width: w / 6 * 5, height: h, title: "Files"},
				},
				{
					Label:             "Databases",
					PleskMenuItemType: MenuItemTypes.ModelType,
					model:             &PlaceHolderModel{width: w / 6 * 5, height: h, title: "Databases"},
				},
				{
					Label:             "Mail",
					PleskMenuItemType: MenuItemTypes.ModelType,
					model:             &PlaceHolderModel{width: w / 6 * 5, height: h, title: "Mail"},
				},
				{
					Label:             "Tools & Settings",
					PleskMenuItemType: MenuItemTypes.ModelType,
					model:             &PlaceHolderModel{width: w / 6 * 5, height: h, title: "Tools & Settings"},
				},
				{
					Label:             "Exit",
					PleskMenuItemType: MenuItemTypes.CommandType,
					command:           tea.Quit,
				},
			},
			MenuFocused: true,
			width:       w / 6,
			height:      h,
		},
		width:  w,
		height: h,
	}
}
