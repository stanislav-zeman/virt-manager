package libvirt

import (
	"fmt"

	"libvirt.org/go/libvirt"
)

const defaultURI = "qemu:///session"

// Libvirt is a facade for interacting with the Libvirt API.
type Libvirt struct {
	conn *libvirt.Connect
}

func New(URI string) (lv Libvirt, err error) {
	if URI == "" {
		URI = defaultURI
	}

	conn, err := libvirt.NewConnect(URI)
	if err != nil {
		err = fmt.Errorf("could not connect to hypervisor: %w", err)
		return
	}

	lv = Libvirt{
		conn: conn,
	}
	return
}
