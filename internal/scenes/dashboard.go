package scenes

import (
	"fmt"

	"github.com/slingercode/pixshell/internal/state"
)

type Dashboard struct{}

func (d *Dashboard) Render(state *state.State) string {
	return fmt.Sprintf("Dashboard\n%s", state.Position.Print())
}
