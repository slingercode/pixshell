package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/slingercode/pixshell/internal/scenes"
	"github.com/slingercode/pixshell/internal/state"
)

var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FFFFFF")).
	Background(lipgloss.Color("#7D56F4")).
	PaddingTop(2).
	PaddingLeft(4).
	Width(22)

type RenderModeEnum int

const (
	Normal RenderModeEnum = iota
	Insert
)

type ScenesEnum int

const (
	Dashboard ScenesEnum = iota
	Editor
	LenghtScenes
)

type Model struct {
	// Render mode of the application
	RenderMode RenderModeEnum
	// Current scene that is being rendered
	Scene ScenesEnum
	// List of the scenes available
	Scenes []scenes.Scene
	// This is the state of the scenes (if that makes sense)
	State *state.State
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.RenderMode {
	case Normal:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "i":
				m.RenderMode = Insert

				return m, nil

			case " ":
				m.RenderMode = Insert
				m.Scene = Editor

				m.State.Position.X++

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
				m.RenderMode = Normal
				m.Scene = Dashboard

				return m, nil
			}
		}
	}

	return m, nil
}

func (m *Model) View() string {
	switch m.Scene {
	case Dashboard:
		return m.Scenes[Dashboard].Render(m.State)

	case Editor:
		return m.Scenes[Editor].Render(m.State)
	}

	return ""
}
