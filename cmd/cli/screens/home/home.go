package home

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type HomeScreen struct {
	width  int
	height int
}

func (hs HomeScreen) Init() tea.Cmd {
	return nil
}

func (hs HomeScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		hs.width = msg.Width
		hs.height = msg.Height
	}
	return hs, nil
}

func (hs HomeScreen) View() string {
	return hs.getBox("hey!")
}

func NewHomeScreen() HomeScreen {
	return HomeScreen{}
}

func (hs HomeScreen) getBox(content string) string {
	box := lipgloss.NewStyle().
		Background(lipgloss.Color("#17151f")).
		Foreground(lipgloss.Color("#F0F0F0")).
		Width(hs.width).Height(hs.height).Padding(5)
	return box.Render(content)
}
