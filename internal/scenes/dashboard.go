package scenes

import (
	"fmt"

	"github.com/slingercode/pixshell/internal/state"
)

type Dashboard struct{}

func (d Dashboard) Render(s state.State) string {
	return fmt.Sprintf("Dashboard\n%s", s.Position.Print())
}
