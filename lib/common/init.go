package common

import (
	"github.com/ProductionPanic/go-cursor"
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
