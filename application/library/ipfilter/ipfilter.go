package ipfilter

import (
	"net/netip"

	"github.com/webx-top/echo"

	"github.com/coscms/webcore/library/ipfilter"
	"github.com/nging-plugins/ftpmanager/application/dbschema"
)

var Factory = ipfilter.NewFactory[uint, uint]().SetGroupGetter(groupGetter)

func IsAllowed(ctx echo.Context, user *dbschema.NgingFtpUser, ip netip.Addr) bool {
	return Factory.IsAllowed(ctx, user.Id, user.GroupId, user.IpBlacklist, user.IpWhitelist, ip)
}

func Validate(ctx echo.Context, ip string) error {
	return ipfilter.ValidateRows(ctx, ip)
}

func groupGetter(ctx echo.Context, groupID uint) (string, string, error) {
	m := dbschema.NewNgingFtpUserGroup(ctx)
	m.Get(nil, `id`, groupID)
	return m.IpBlacklist, m.IpWhitelist, nil
}
