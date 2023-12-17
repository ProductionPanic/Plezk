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

type BackMsg struct{}

func Back() tea.Cmd {
	return func() tea.Msg {
		return BackMsg{}
	}
}
