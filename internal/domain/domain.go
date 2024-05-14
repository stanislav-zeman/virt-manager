package domain

import (
	"libvirt.org/go/libvirt"
)

type State string

const (
	NoState        State = "No state"
	RunningState   State = "Running"
	BlockedState   State = "Blocked"
	PausedState    State = "Paused"
	ShutdownState  State = "Shutdown"
	CrashedState   State = "Crashed"
	SuspendedState State = "Suspended"
	ShutoffState   State = "Shutoff"
)

type Domain struct {
	Name  string `json:"name"`
	State State  `json:"status"`
}

func New(domain libvirt.Domain) (Domain, error) {
	name, err := domain.GetName()
	if err != nil {
		return Domain{}, err
	}

	libvirtStatus, _, err := domain.GetState()
	if err != nil {
		return Domain{}, err
	}

	var state State
	switch libvirtStatus {
	case libvirt.DOMAIN_NOSTATE:
		state = NoState
	case libvirt.DOMAIN_RUNNING:
		state = RunningState
	case libvirt.DOMAIN_BLOCKED:
		state = BlockedState
	case libvirt.DOMAIN_PAUSED:
		state = PausedState
	case libvirt.DOMAIN_SHUTDOWN:
		state = ShutdownState
	case libvirt.DOMAIN_CRASHED:
		state = CrashedState
	case libvirt.DOMAIN_PMSUSPENDED:
		state = SuspendedState
	case libvirt.DOMAIN_SHUTOFF:
		state = ShutoffState
	}

	return Domain{
		Name:  name,
		State: state,
	}, nil
}
