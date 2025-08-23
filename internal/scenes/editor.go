package scenes

import (
	"fmt"

	"github.com/slingercode/pixshell/internal/state"
)

type Editor struct{}

func (e *Editor) Render(state *state.State) string {
	return fmt.Sprintf("Editor\n%s", state.Position.Print())
}
