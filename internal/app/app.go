package app

import (
	"context"
	"fmt"
)

type App struct {
	ctx context.Context
}

func New() *App {
	return &App{}
}

// Startup is called when the app starts.
// The context is saved so we can call the runtime methods.
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name.
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
