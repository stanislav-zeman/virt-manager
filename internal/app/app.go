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

func (a *App) Close(ctx context.Context) {
	a.lv.Close()
}

func (a *App) Domains() ([]domain.Domain, error) {
	a.log.Info("fetching domains")
	if a.lv == nil {
		return nil, errors.New("no connection set up")
	}

	return a.lv.Domains()
}

func (a *App) Connect(URI string) error {
	if URI == "" {
		err := errors.New("blank URI provided")
		a.log.Error("invalid URI",
			zap.Error(err),
		)
		return err
	}

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

func (a *App) Connected() bool {
	connected := a.lv != nil
	a.log.Info("retrieved connection status",
		zap.Bool("connected", connected),
	)
	return connected
}

func (a *App) Create(domainName string, vCPUs uint, memory uint64) error {
	return a.lv.Create(domainName, vCPUs, memory)
}

func (a *App) Suspend(domainName string) error {
	return a.lv.Suspend(domainName)
}

func (a *App) Resume(domainName string) error {
	return a.lv.Resume(domainName)
}

func (a *App) Shutdown(domainName string) error {
	return a.lv.Destroy(domainName)
}

func (a *App) Destroy(domainName string) error {
	return a.lv.Destroy(domainName)
}
