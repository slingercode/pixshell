package internal

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FFFFFF")).
	Background(lipgloss.Color("#7D56F4")).
	PaddingTop(2).
	PaddingLeft(4).
	Width(22)

type ModelState int

const (
	Normal ModelState = iota
	Insert
)

type ModelScene int

const (
	Dashboard ModelScene = iota
	Editor
)

type Model struct {
	// State of the application
	State ModelState
	// Current scene that is being rendered
	Scene ModelScene
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.State {
	case Normal:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "i":
				m.State = Insert
				return m, nil
			case " ":
				m.State = Insert
				m.Scene = Editor
				return m, nil
			case "ctrl+c", "q":
				return m, tea.Quit
			}
		}

	case Insert:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "q":
				m.State = Normal
				m.Scene = Dashboard
				return m, nil
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	switch m.Scene {
	case Dashboard:
		return style.Render("Hello World")
	case Editor:
		return "Yes"
	}

	return ""
}
