package ftpmanager

import (
	"github.com/coscms/webcore/registry/navigate"

	"github.com/nging-plugins/ftpmanager/application/handler"
)

var LeftNavigate = handler.LeftNavigate

func RegisterNavigate(nc *navigate.Collection) {
	nc.Backend.AddLeftItems(-1, LeftNavigate)
}
