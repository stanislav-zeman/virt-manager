package libvirt

import (
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

func (f *Facade) Create(domainName string, vCPUs uint, memory uint64) error {
	domain := libvirt.Domain{}
	err := domain.Rename(domainName, 0)
	if err != nil {
		return err
	}

	err = domain.SetVcpus(vCPUs)
	if err != nil {
		return err
	}

	err = domain.SetMemory(memory)
	if err != nil {
		return err
	}

	flags := libvirt.DomainCreateFlags(0)
	err = domain.CreateWithFlags(flags)
	if err != nil {
		return err
	}

	domain.Free()
	return nil
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
