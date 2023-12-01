package websites

import (
	"fmt"
	"github.com/ProductionPanic/go-pretty"
	tea "github.com/charmbracelet/bubbletea"
	"plezk/lib/domain"
)

type BubbleTeaWebsiteModule struct {
	domainsWithChildDomains []domain.Domain
	cursor                  int
	selectedDomain          *domain.Domain
}

func (m *BubbleTeaWebsiteModule) Init() tea.Cmd {
	return nil
}

func (m *BubbleTeaWebsiteModule) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.domainsWithChildDomains)-1 {
				m.cursor++
			}
		case "enter":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *BubbleTeaWebsiteModule) View() string {
	s := "  [blue,bold]Plezk:[]\n"
	for i, item := range m.domainsWithChildDomains {
		cursor := " "
		if m.cursor == i {
			cursor = "[cyan]>[]"
		}
		s += fmt.Sprintf("%s %s\n", cursor, item)
	}

	return pretty.Parse(s)
}

func RenderBubbleTeaWebsiteModule(domainsWithChildDomains []domain.Domain) string {
	p := tea.NewProgram(&BubbleTeaWebsiteModule{
		domainsWithChildDomains: domainsWithChildDomains,
		cursor:                  0,
		selectedDomain:          nil,
	})
	m, _ := p.Run()
	return domainsWithChildDomains[m.(*BubbleTeaWebsiteModule).cursor].Name
}

func Start() {
	domains := domain.List()
	RenderBubbleTeaWebsiteModule(domains)

}
