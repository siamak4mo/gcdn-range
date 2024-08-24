package providers

import (
	"errors"
)

type CDN_Provider interface {
	GET() ([]string, error)
}

type Provider struct {
	Name  string
	Pr    CDN_Provider
	DLerr error
	CIDR  []string
	id    cdn_type
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

func newProvider(pr CDN_Provider, name string, idx cdn_type) Provider {
	return Provider{
		Pr:    pr,
		Name:  name,
		id:    idx,
		CIDR:  make([]string, 0),
		DLerr: nil,
	}
}

// index of Provider in this array must be equal to it's ID
var CDNs = []Provider{
	newProvider(nil,
		"maxcdn", Maxcdn_CDN),
	newProvider(CloudFlare__P{},
		"cloudflare", Cloudflare_CDN),
	newProvider(nil,
		"fastly", Fastly_CDN),
	newProvider(nil,
		"incapsula", Incapsula_CDN),
	newProvider(nil,
		"cachefly", Cachefly_CDN),
	newProvider(nil,
		"cloudfront", Cloudfront_CDN),
}

func (p *Provider) DoFetch() *Provider {
	if p == nil || p.Pr == nil {
		if p != nil {
			p.DLerr = errors.New("Not Implemented")
		}
		return p
	}
	s, e := p.Pr.GET()
	if e != nil {
		p.DLerr = e
		return p
	}

	p.CIDR = s
	p.DLerr = e
	return p
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

func MkCDN(names []string) []*Provider {
	r := make([]*Provider, 0)
	for _, name := range names {
		s, e := SearchCDN(name)
		if e != nil {
			r = append(r, &s)
		}
	}
	return r
}

func MKCDN_all() []*Provider {
	r := make([]*Provider, 0)
	for _, pcpy := range CDNs {
		r = append(r, &pcpy)
	}
	return r
}
