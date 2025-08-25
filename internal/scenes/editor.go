package scenes

import (
	"github.com/slingercode/pixshell/internal/components"
	"github.com/slingercode/pixshell/internal/state"
)

type Editor struct{}

func (e Editor) Render(s state.State) string {
	return components.GridComponent(s)
}
