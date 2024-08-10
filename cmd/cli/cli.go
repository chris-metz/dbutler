package main

import (
	"github.com/chris-metz/dbutler/lib/db"
)

func main() {
	dbHandler := db.NewDbHandler()
	defer dbHandler.Shutdown()
	dbHandler.ReCreateSchema()
	dbHandler.SeedDatabase()
	// rs := root.NewRootScreen()
	// p := tea.NewProgram(rs)
	// _, err := p.Run()
	// if err != nil {
	// 	panic(err)
	// }
}
