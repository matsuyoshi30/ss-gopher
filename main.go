package main

import (
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const refreshInterval = 500 * time.Millisecond

var (
	view *tview.Box
	app  *tview.Application

	flag = true
)

type Pic []string

func (p Pic) width() int {
	return len(p[0])
}

func (p Pic) height() int {
	return len(p)
}

type Picmap map[int]Pic

func draw(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
	var gopher Pic

	if flag {
		gopher = frontpage[1]
	} else {
		gopher = frontpage[2]
	}

	startY := (height - gopher.height()) / 2
	for _, gopherStr := range gopher {
		tview.Print(screen, gopherStr, x, startY, width, tview.AlignCenter, tcell.ColorWhite)
		startY++
	}

	flag = !flag

	return 0, 0, 0, 0
}

func run() error {
	app = tview.NewApplication()
	view = tview.NewBox().SetDrawFunc(draw)

	return app.SetRoot(view, true).Run()
}

func refresh() {
	tick := time.NewTicker(refreshInterval)
	for {
		select {
		case <-tick.C:
			app.Draw()
		}
	}
}

func main() {
	go refresh()
	if err := run(); err != nil {
		panic(err)
	}
}
