package libvirt

import (
	"encoding/xml"
	"fmt"

	"github.com/stanislav-zeman/virt-manager/internal/domain"
	"libvirt.org/go/libvirt"
)

// Facade is a facade for interacting with the Libvirt API.
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

func (f *Facade) Close() {
	f.conn.Close()
}

func (f *Facade) Domains() ([]domain.Domain, error) {
	flags := libvirt.ConnectListAllDomainsFlags(0)
	domains, err := f.conn.ListAllDomains(flags)
	if err != nil {
		return nil, err
	}

	apiDomains := make([]domain.Domain, 0, len(domains))
	for _, d := range domains {
		apiDomain, err := domain.New(d)
		if err != nil {
			return nil, err
		}

		d.Free()
		apiDomains = append(apiDomains, apiDomain)
	}

	return apiDomains, nil
}

func (f *Facade) Start(domainName string) error {
	domain, err := f.conn.LookupDomainByName(domainName)
	if err != nil {
		return err
	}

	defer domain.Free()
	return domain.Create()
}

func (f *Facade) Resume(domainName string) error {
	domain, err := f.conn.LookupDomainByName(domainName)
	if err != nil {
		return err
	}

	defer domain.Free()
	return domain.Resume()
}

func (f *Facade) Suspend(domainName string) error {
	domain, err := f.conn.LookupDomainByName(domainName)
	if err != nil {
		return err
	}

	defer domain.Free()
	return domain.Suspend()
}

func (f *Facade) Shutdown(domainName string) error {
	domain, err := f.conn.LookupDomainByName(domainName)
	if err != nil {
		return err
	}

	defer domain.Free()
	return domain.Shutdown()
}

func (f *Facade) Destroy(domainName string) error {
	domain, err := f.conn.LookupDomainByName(domainName)
	if err != nil {
		return err
	}

	defer domain.Free()
	return domain.Destroy()
}

func (f *Facade) AttachDevice(domainName, diskType, diskDevice, sourcePath, targetDevice string) error {
	d, err := f.conn.LookupDomainByName(domainName)
	if err != nil {
		return err
	}

	defer d.Free()
	device := domain.Disk{
		Type:   diskType,
		Device: diskDevice,
		Driver: domain.Driver{
			Name:  "qemu",
			Type:  "raw",
			Cache: "none",
		},
		Source: domain.Source{
			File: sourcePath,
		},
		Target: domain.Target{
			Dev: targetDevice,
		},
	}
	bytes, err := xml.Marshal(device)
	if err != nil {
		return err
	}

	err = d.AttachDevice(string(bytes))
	if err != nil {
		return err
	}

	return nil

}

func (f *Facade) DetachDevice(domainName, diskType, diskDevice, sourcePath, targetDevice string) error {
	d, err := f.conn.LookupDomainByName(domainName)
	if err != nil {
		return err
	}

	defer d.Free()
	device := domain.Disk{
		Type:   diskType,
		Device: diskDevice,
		Driver: domain.Driver{
			Name:  "qemu",
			Type:  "raw",
			Cache: "none",
		},
		Source: domain.Source{
			File: sourcePath,
		},
		Target: domain.Target{
			Dev: targetDevice,
		},
	}
	bytes, err := xml.Marshal(device)
	if err != nil {
		return err
	}

	err = d.DetachDevice(string(bytes))
	if err != nil {
		return err
	}

	return nil
}
