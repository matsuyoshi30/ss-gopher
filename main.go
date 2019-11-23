package main

import (
	"time"

	"github.com/rivo/tview"
)

const refreshInterval = 500 * time.Millisecond

var (
	view *tview.Box
	app  *tview.Application
)

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
