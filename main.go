package main

import (
	"context"
	"log"

	"github.com/slingercode/pixshell/internal/cmd"
)

func main() {
	ctx := context.Background()

	if err := cmd.Init(ctx); err != nil {
		log.Fatal(err)
	}
}
