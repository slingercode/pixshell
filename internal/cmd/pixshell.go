package cmd

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/slingercode/pixshell/internal"
)

func initialModel() internal.Model {
	return internal.Model{
		State: internal.Normal,
		Scene: internal.Dashboard,
	}
}

func Init(ctx context.Context) error {
	p := tea.NewProgram(
		initialModel(),
		tea.WithContext(ctx),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}
