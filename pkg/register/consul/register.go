package consul

import (
	"fmt"
	"github.com/dotamixer/doom/pkg/net/ip"
	"github.com/dotamixer/doom/pkg/register"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
)

type Register struct {
	agent   *api.Agent
	id      string
	opt     *register.Options
	localIp string
}

func NewRegister(opt *register.Options) (*Register, error) {
	r := &Register{
		opt: opt,
	}

	config := api.DefaultConfig()
	config.Address = opt.RegistryAddr

	client, err := api.NewClient(config)
	if err != nil {
		logrus.Errorf("Failed to new consul client. err:[%v]", err)
		return nil, err
	}

	r.localIp = ip.InternalIP()
	r.agent = client.Agent()
	r.id = fmt.Sprintf("%s-%s-%d", opt.Name, r.localIp, opt.Port)

	return r, nil
}

func (r *Register) Register() (err error) {

	reg := &api.AgentServiceRegistration{
		Kind:              "",
		ID:                "",
		Name:              r.opt.Name,
		Tags:              nil,
		Port:              r.opt.Port,
		Address:           r.localIp,
		TaggedAddresses:   nil,
		EnableTagOverride: false,
		Meta:              nil,
		Weights:           nil,
		Check:             nil,
		Checks:            nil,
		Proxy:             nil,
		Connect:           nil,
		Namespace:         "",
	}

	if err = r.agent.ServiceRegister(reg); err != nil {
		logrus.Fatalf("Failed to register service. err:[%s]", err)
		return err
	}

	return
}

func (r *Register) Deregister() (err error) {
	err = r.agent.ServiceDeregister(r.id)
	if err != nil {
		logrus.Fatalf("Failed to deregister service. err:[%v]", err)
	}

	return err
}
