package providers

import (
	"errors"
)

type CDN_Provider interface {
	GET(results []string) ([]string, error)
}

type Provider struct {
	Pr   CDN_Provider
	Name string
	ID   cdn_type
}

type cdn_type int

const (
	Maxcdn_CDN = iota
	Cloudflare_CDN
	Fastly_CDN
	Incapsula_CDN
	Cachefly_CDN
	Cloudfront_CDN
)

const (
	CDN_NOT_FOUND_ERR      = "CDN not found"
	CDN_NAME_NOT_FOUND_ERR = "CDN name not found"
)

// index of Provider in this array must be equal to it's ID
var CDNs = []Provider{
	{nil,
		"maxcdn", Maxcdn_CDN},
	{CloudFlare__P{},
		"cloudflare", Cloudflare_CDN},
	{nil,
		"fastly", Fastly_CDN},
	{nil,
		"incapsula", Incapsula_CDN},
	{nil,
		"cachefly", Cachefly_CDN},
	{nil,
		"cloudfront", Cloudfront_CDN},
}

func GetCDN(prov int) (Provider, error) {
	if prov >= len(CDNs) {
		return Provider{}, errors.New(CDN_NOT_FOUND_ERR)
	}
	return CDNs[prov], nil
}

func SearchCDN(name string) (Provider, error) {
	switch name {
	case "cloudflare":
		return CDNs[Cloudflare_CDN], nil

	default:
		return Provider{}, errors.New(CDN_NAME_NOT_FOUND_ERR)
	}
}

func MkCDN(names []string) []Provider {
	r := make([]Provider, 0)
	for _, name := range names {
		s, e := SearchCDN(name)
		if e != nil {
			r = append(r, s)
		}
	}
	return r
}
