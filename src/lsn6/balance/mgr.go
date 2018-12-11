package balance

import (
	"fmt"

	"github.com/luci/go-render/render"
)

type BalancerMgr struct {
	allBalancer map[string]Balancer
}

var mgr = BalancerMgr{
	allBalancer: make(map[string]Balancer),
}

func (p *BalancerMgr) RegisterBalancer(name string, b Balancer) {
	p.allBalancer[name] = b
}

func RegisterBalancer(name string, b Balancer) {
	mgr.RegisterBalancer(name, b)
}

func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	if balancer, ok := mgr.allBalancer[name]; !ok {
		err = fmt.Errorf("Not dound %s balancer", name)
	} else {
		inst, err = balancer.DoBalance(insts)
		fmt.Println("balancer type ", render.Render(balancer))
	}
	return
}
