package common

import "github.com/ProductionPanic/go-cursor"

func Cls() {
	cursor.ClearScreen()
	cursor.Top()
}
