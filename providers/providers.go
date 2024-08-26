package providers

import (
	"bufio"
	"errors"
	"net/http"
)

type ProvChan chan string

type CDN_Provider interface {
	GET(cout ProvChan, flags int) error
}

// flags
const (
	DL_IPv4 = iota
	DL_IPv6
	DL_ALL
)

type Provider struct {
	Pr    CDN_Provider
	CIDR  ProvChan
	DLerr error
	Name  string
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
	Arvan_CDN
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
		DLerr: nil,
	}
}

// index of Provider in this array must be equal to it's ID
var CDNs = []Provider{
	newProvider(Maxcdn__P{},
		"maxcdn", Maxcdn_CDN),
	newProvider(CloudFlare__P{},
		"cloudflare", Cloudflare_CDN),
	newProvider(Fastly__P{},
		"fastly", Fastly_CDN),
	newProvider(Incapsula__P{},
		"incapsula", Incapsula_CDN),
	newProvider(Cachefly__P{},
		"cachefly", Cachefly_CDN),
	newProvider(CloudFront__P{},
		"cloudfront", Cloudfront_CDN),
	newProvider(Arvan__P{},
		"arvancloud", Arvan_CDN),
}

func (p *Provider) DoFetch(flags int) *Provider {
	if p == nil || p.Pr == nil {
		if p != nil {
			p.DLerr = errors.New("Not Implemented")
		}
		return p
	}
	e := p.Pr.GET(p.CIDR, flags)
	if e != nil {
		p.DLerr = e
		return p
	}
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
	case "cloudflare", "cf":
		return CDNs[Cloudflare_CDN], nil

	case "max", "maxcdn":
		return CDNs[Maxcdn_CDN], nil

	case "fastly", "fastlycdn", "fastlycloud":
		return CDNs[Fastly_CDN], nil

	case "incapsula", "incapsulacdn":
		return CDNs[Incapsula_CDN], nil

	case "cachefly":
		return CDNs[Cachefly_CDN], nil

	case "cloudfront", "aws", "amazon", "amazonaws":
		return CDNs[Cloudfront_CDN], nil

	case "arvan", "arvancloud":
		return CDNs[Arvan_CDN], nil
	default:
		return Provider{}, errors.New(CDN_NAME_NOT_FOUND_ERR)
	}
}

func MkProv(names []string) []*Provider {
	r := make([]*Provider, 0)
	for _, name := range names {
		s, e := SearchCDN(name)
		if e == nil {
			r = append(r, &s)
		}
	}
	return r
}

func MKProvs() []*Provider {
	r := make([]*Provider, 0)
	for i, _ := range CDNs {
		r = append(r, &CDNs[i])
	}
	return r
}

// internal functions
func rw2cahn_from_http(res *http.Response, cout chan string) {
	scanner := bufio.NewScanner(res.Body)
	var e error
	for ; scanner.Scan() && e == nil; e = scanner.Err() {
		cout <- scanner.Text()
	}
}
