package website

import (
	"fmt"
	"github.com/ProductionPanic/go-pretty"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
	"plezk/lib/common"
	"plezk/lib/domain"
)

type (
	errMsg error
)

var inputs []textinput.Model = make([]textinput.Model, 3)

func Add() int {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	success := domain.Create(
		inputs[dom].Value(),
		inputs[ftplogin].Value(),
		inputs[ftppass].Value(),
	)

	common.Cls()
	if success {
		pretty.Print("[green]Successfully added website[]")
	} else {
		pretty.Print("[red]Failed to add website[]")
	}

	return 0
}

const (
	dom = iota
	ftplogin
	ftppass
)

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

var (
	inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
)

type model struct {
	inputs  []textinput.Model
	focused int
	err     error
}

func initialModel() model {

	inputs[dom] = textinput.New()
	inputs[dom].Prompt = "Domain name: "
	inputs[dom].Focus()
	inputs[dom].Width = 30
	inputs[dom].CharLimit = 60

	inputs[ftplogin] = textinput.New()
	inputs[ftplogin].Prompt = "FTP Login: "
	inputs[ftplogin].Width = 30
	inputs[ftplogin].CharLimit = 60

	inputs[ftppass] = textinput.New()
	inputs[ftppass].Prompt = "FTP Password: "
	inputs[ftppass].Width = 30
	inputs[ftppass].CharLimit = 60
	inputs[ftppass].EchoCharacter = '*'
	inputs[ftppass].EchoMode = textinput.EchoPassword

	return model{
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				return m, tea.Quit
			}
			m.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return fmt.Sprintf(
		` %s
 %s

 %s  
 %s  %s

 %s
`,
		inputStyle.Width(30).Render("Domain: "),
		m.inputs[dom].View(),
		inputStyle.Width(30).Render("Login details: "),
		m.inputs[ftplogin].View(),
		m.inputs[ftppass].View(),
		continueStyle.Render("Continue ->"),
	) + "\n"
}

// nextInput focuses the next input field
func (m *model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *model) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
