package main

import (
	"golang.org/x/sys/unix"
)

func GetTermSize() (int, int) {
	ws, err := unix.IoctlGetWinsize(int(unix.Stdout), unix.TIOCGWINSZ)
	if err != nil {
		panic(err)
	}
	return int(ws.Col), int(ws.Row)
}
