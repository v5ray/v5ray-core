package geodata

import (
	"github.com/v4fly/v4ray-core/v0/app/router/routercommon"
)

//go:generate go run github.com/v4fly/v4ray-core/v0/common/errors/errorgen

type LoaderImplementation interface {
	LoadSite(filename, list string) ([]*routercommon.Domain, error)
	LoadIP(filename, country string) ([]*routercommon.CIDR, error)
}

type Loader interface {
	LoaderImplementation
	LoadGeoSite(list string) ([]*routercommon.Domain, error)
	LoadGeoSiteWithAttr(file string, siteWithAttr string) ([]*routercommon.Domain, error)
	LoadGeoIP(country string) ([]*routercommon.CIDR, error)
}
