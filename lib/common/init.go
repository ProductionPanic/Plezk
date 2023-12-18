package common

import (
	"github.com/ProductionPanic/go-cursor"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/eiannone/keyboard"
)

func Cls() {
	cursor.ClearScreen()
	cursor.Top()
}

func Pause() {
	_, _, err := keyboard.GetKey()
	if err != nil {
		return
	}
}

type BackToMenuMsg struct{}

var BackToMenu = tea.Cmd(func() tea.Msg {
	return BackToMenuMsg{}
})

func SelectModel(model string) tea.Cmd {
	return func() tea.Msg {
		return MenuSelectMsg{model}
	}
}

type MenuSelectMsg struct {
	Model string
}
