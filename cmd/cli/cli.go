package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/chris-metz/dbutler/cmd/cli/screens/home"
	"github.com/chris-metz/dbutler/lib/db"
)

func main() {
	dbHandler := db.NewDbHandler()
	defer dbHandler.Shutdown()
	dbHandler.ReCreateSchema()
	dbHandler.SeedDatabase()
	hs := home.NewHomeScreen()
	p := tea.NewProgram(hs)
	_, err := p.Run()
	if err != nil {
		panic(err)
	}
}
