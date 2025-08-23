package cmd

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	model "github.com/slingercode/pixshell/internal"
	"github.com/slingercode/pixshell/internal/scenes"
	"github.com/slingercode/pixshell/internal/state"
	"github.com/slingercode/pixshell/internal/vec4"
)

func initialModel() model.Model {
	dashboard := scenes.Dashboard{}
	editor := scenes.Editor{}

	scenes := make([]scenes.Scene, model.LenghtScenes)
	scenes[model.Dashboard] = &dashboard
	scenes[model.Editor] = &editor

	state := state.State{
		Position: vec4.Vec4{X: 0.0, Y: 0.0, Z: 0.0, A: 0.0},
	}

	return model.Model{
		RenderMode: model.Normal,
		Scene:      model.Dashboard,
		Scenes:     scenes,
		State:      &state,
	}
}

func Init(ctx context.Context) error {
	model := initialModel()

	p := tea.NewProgram(
		&model,
		tea.WithContext(ctx),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}
