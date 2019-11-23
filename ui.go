package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Pic []string

func (p Pic) width() int {
	return len(p[0])
}

func (p Pic) height() int {
	return len(p)
}

type Picmap map[int]Pic

var flag = true

func draw(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
	var gopher Pic

	if flag {
		gopher = appenginegopher[1]
	} else {
		gopher = appenginegopher[2]
	}

	startY := (height - gopher.height()) / 2
	for _, gopherStr := range gopher {
		tview.Print(screen, gopherStr, x, startY, width, tview.AlignCenter, tcell.ColorWhite)
		startY++
	}

	flag = !flag

	return 0, 0, 0, 0
}
