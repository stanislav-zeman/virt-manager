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

func (a *App) Start(domainName string) error {
	err := a.lv.Start(domainName)
	if err != nil {
		a.log.Error("failed starting domain",
			zap.Error(err),
			zap.String("name", domainName),
		)
		return err
	}

	a.log.Info("started domain",
		zap.String("name", domainName),
	)
	return nil

}

func (a *App) Suspend(domainName string) error {
	err := a.lv.Suspend(domainName)
	if err != nil {
		a.log.Error("failed suspending domain",
			zap.Error(err),
			zap.String("name", domainName),
		)
		return err
	}

	a.log.Info("suspended domain",
		zap.String("name", domainName),
	)
	return nil
}

func (a *App) Resume(domainName string) error {
	err := a.lv.Resume(domainName)
	if err != nil {
		a.log.Error("failed resuming domain",
			zap.Error(err),
			zap.String("name", domainName),
		)
		return err
	}

	a.log.Info("resumed domain",
		zap.String("name", domainName),
	)
	return nil
}

func (a *App) Shutdown(domainName string) error {
	err := a.lv.Shutdown(domainName)
	if err != nil {
		a.log.Error("failed shutdowning domain",
			zap.Error(err),
			zap.String("name", domainName),
		)
		return err
	}

	a.log.Info("shutdowned domain",
		zap.String("name", domainName),
	)
	return nil
}

func (a *App) Destroy(domainName string) error {
	err := a.lv.Destroy(domainName)
	if err != nil {
		a.log.Error("failed destroying domain",
			zap.Error(err),
			zap.String("name", domainName),
		)
		return err
	}

	a.log.Info("destroyed domain",
		zap.String("name", domainName),
	)
	return nil
}

func (a *App) AttachDevice(domainName, diskType, diskDevice, path, target string) error {
	err := a.lv.AttachDevice(domainName, diskType, diskDevice, path, target)
	if err != nil {
		a.log.Error("failed attaching device",
			zap.Error(err),
			zap.String("name", domainName),
			zap.String("path", path),
		)
		return err
	}

	a.log.Info("attached device",
		zap.String("name", domainName),
		zap.String("path", path),
	)
	return nil
}

func (a *App) DetachDevice(domainName, diskType, diskDevice, path, target string) error {
	err := a.lv.DetachDevice(domainName, diskType, diskDevice, path, target)
	if err != nil {
		a.log.Error("failed detaching device",
			zap.Error(err),
			zap.String("name", domainName),
			zap.String("path", path),
		)
		return err
	}

	a.log.Info("detached device",
		zap.String("name", domainName),
		zap.String("path", path),
	)
	return nil
}
