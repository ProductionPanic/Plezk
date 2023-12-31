package websites

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	"plezk/lib/colors"
	"plezk/lib/common"
	"plezk/lib/domain"
)

type DomainsAndWebsitesModel struct {
	Width          int
	Height         int
	domains        []domain.Domain
	cursor         int
	SelectedDomain int
	focused        bool
	DomainModel    tea.Model
}

const (
	treePipe  = "│"
	treeTPipe = "├"
	treeLPipe = "└"
)

func (m DomainsAndWebsitesModel) Init() tea.Cmd {
	m.domains = domain.List()
	m.cursor = 0
	m.SelectedDomain = -1
	m.focused = true
	return nil
}

func (m *DomainsAndWebsitesModel) TotalLength() int {
	total := 0
	for _, dom := range m.domains {
		total += dom.Count()
	}
	return total
}

func (m *DomainsAndWebsitesModel) GetSelectedDomain() *domain.Domain {
	index := m.SelectedDomain
	for _, dom := range m.domains {
		if index == 0 {
			return &dom
		}
		index--
		if len(dom.Children) > 0 {
			for _, child := range dom.Children {
				if index == 0 {
					return &child
				}
				index--
			}
		}
	}
	return nil
}

func DomainSelect(domain domain.Domain) tea.Cmd {
	return func() tea.Msg {
		return DomainSelectMsg{domain}
	}
}

type DomainSelectMsg struct {
	Domain domain.Domain
}

func (m *DomainsAndWebsitesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	if !m.focused && m.DomainModel != nil {
		var cmd tea.Cmd
		m.DomainModel, cmd = m.DomainModel.Update(msg)
		cmds = append(cmds, cmd)
	}
	if m.TotalLength() == 0 {
		m.domains = domain.List()
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if m.cursor < m.TotalLength()-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = m.TotalLength() - 1
			}
		case "enter", "right", "l":
			m.SelectedDomain = m.cursor
			m.focused = false
			return m, DomainSelect(*m.GetSelectedDomain())
		case "esc", "left", "h":
			return m, common.BackToMenu
		}

	}

	return m, nil
}

func (m DomainsAndWebsitesModel) View() string {
	if !m.focused && m.DomainModel != nil {
		return m.DomainModel.View()
	}
	rootStyle := lg.NewStyle().Background(lg.Color("#000000"))
	domainStyle := lg.NewStyle().Background(colors.Black).
		Border(lg.DoubleBorder(), false, false, false, true).Foreground(colors.White).
		BorderForeground(colors.Black).BorderBackground(colors.Black).Width(m.Width - 1)

	var domains []string
	totali := 0
	for _, dom := range m.domains {
		domstyle := domainStyle.Copy()
		if totali == m.cursor {
			domstyle = domstyle.BorderForeground(colors.Secondary)
		}
		if totali == m.SelectedDomain {
			domstyle = domstyle.Foreground(colors.Secondary)
		}
		totali++
		domains = append(domains, domstyle.Render(dom.Name))
		for ci, child := range dom.Children {
			pipechar := fmt.Sprintf(" %s ", treeTPipe)
			if ci == len(dom.Children)-1 {
				pipechar = fmt.Sprintf(" %s ", treeLPipe)
			}
			if totali == m.cursor {
				domains = append(domains, domainStyle.Copy().BorderForeground(colors.Secondary).PaddingLeft(1).Render(pipechar+child.Name))
				totali++
				continue
			}
			domains = append(domains, domainStyle.Copy().PaddingLeft(1).Render(pipechar+child.Name))
			totali++
		}
	}

	dom := lg.JoinVertical(
		lg.Left,
		domains...,
	)
	dom = lg.NewStyle().Background(colors.Black).Render(dom)
	headr := lg.NewStyle().
		Bold(true).
		Foreground(colors.Secondary).
		Background(colors.Black).
		PaddingLeft(1).
		Width(m.Width).Height(2).Render("Websites and domains")

	return rootStyle.Render(lg.JoinVertical(lg.Left, headr, dom))
}
