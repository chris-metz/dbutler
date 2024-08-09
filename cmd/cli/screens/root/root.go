package root

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/chris-metz/dbutler/cmd/cli/screens/home"
	selectconnection "github.com/chris-metz/dbutler/cmd/cli/screens/select_connection"
)

type RootScreen struct {
	state            string
	selectConnScreen selectconnection.SelectConnectionScreen
	homeScreen       home.HomeScreen
}

func (rs RootScreen) Init() tea.Cmd {
	return nil
}

func (rs RootScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return rs, tea.Quit
		case "j":
			if rs.state == "root" {
				rs.state = "home"
			} else {
				rs.state = "root"
			}
			return rs, nil
		}
		return rs, nil
	}

	var cmd tea.Cmd
	if rs.state == "root" {
		var scModel tea.Model
		scModel, cmd = rs.selectConnScreen.Update(msg)
		rs.selectConnScreen = scModel.(selectconnection.SelectConnectionScreen)
	}
	return rs, cmd
}

func (rs RootScreen) View() string {
	tea.Println("asdf")
	if rs.state == "home" {
		return rs.homeScreen.View()
	} else {
		return rs.selectConnScreen.View()
	}
}

func NewRootScreen() RootScreen {
	return RootScreen{
		state:            "root",
		homeScreen:       home.NewHomeScreen(),
		selectConnScreen: selectconnection.NewSelectConnectionScreen(),
	}
}
