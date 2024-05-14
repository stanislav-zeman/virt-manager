package app

import (
	"context"
	"errors"

	"github.com/stanislav-zeman/virt-manager/internal/domain"
	"github.com/stanislav-zeman/virt-manager/internal/libvirt"
	"go.uber.org/zap"
)

type App struct {
	log *zap.Logger
	ctx context.Context
	lv  *libvirt.Facade
}

func New() *App {
	log, _ := zap.NewProduction()
	defer log.Sync()
	return &App{
		log: log,
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Domains() (domains []domain.Domain, err error) {
	a.log.Info("fetching domains")
	if a.lv == nil {
		err = errors.New("No connection to a hypervisor was made. Did you specify the URI in the settings tab?")
		return
	}

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
	}, nil
}

func (a *App) Connect(URI string) error {
	if URI == "" {
		a.log.Error("blank URI provided")
		return errors.New("Blank URI is invalid!")
	}

	a.log.Info("establishing new connection")
	lv, err := libvirt.New(URI)
	if err != nil {
		a.log.Error("could not connect to hypervisor",
			zap.Error(err),
			zap.String("URI", URI),
		)
		return err
	}

	a.log.Info("connected to hypervisor",
		zap.String("URI", URI),
	)
	a.lv = lv
	return nil
}
