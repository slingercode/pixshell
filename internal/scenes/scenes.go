package scenes

import "github.com/slingercode/pixshell/internal/state"

type Scene interface {
	Render(state state.State) string
}
