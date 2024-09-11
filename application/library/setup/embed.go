package setup

import (
	_ "embed"

	"github.com/coscms/webcore/library/config"
)

//go:embed install.sql
var installSQL string

func RegisterSQL(sc *config.SQLCollection) {
	sc.RegisterInstall(`nging`, installSQL)
}
