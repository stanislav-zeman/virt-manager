package app

import (
	"context"
	"fmt"

	"github.com/stanislav-zeman/virt-manager/internal/domain"
	"github.com/stanislav-zeman/virt-manager/internal/libvirt"
)

type App struct {
	ctx context.Context
	lv  libvirt.Libvirt
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

func (a *App) Domains() []domain.Domain {
	return []domain.Domain{
		{
			Name:   "prettyCoolVM",
			Status: "running",
		},
		{
			Name:   "notSoCoolVM",
			Status: "suspended",
		},
		{
			Name:   "nobodyCaresVM",
			Status: "stopped",
		},
	}
}

func (a *App) Connect(URI string) error {
	lv, err := libvirt.New(URI)
	if err != nil {
		return err
	}

	a.lv = lv
	return nil
}
