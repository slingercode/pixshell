package cmd

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	model "github.com/slingercode/pixshell/internal"
)

func Init(ctx context.Context) error {
	model := model.Init()

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
