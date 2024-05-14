package libvirt

import (
	"fmt"

	"libvirt.org/go/libvirt"
)

// Facade is a facade for interacting with the Facade API.
type Facade struct {
	conn *libvirt.Connect
}

func New(URI string) (lv *Facade, err error) {
	conn, err := libvirt.NewConnect(URI)
	if err != nil {
		err = fmt.Errorf("could not connect to hypervisor: %w", err)
		return
	}

	lv = &Facade{
		conn: conn,
	}
	return
}
