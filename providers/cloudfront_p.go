package providers

import (
	"encoding/json"
	"io"
	"net/http"
)

type CloudFront__P struct {
	RAW      []byte
	Prefixes []struct {
		IP_pref  string `json:"ip_prefix"`
		IP6_pref string `json:"ip6_prefix"`
	} `json:"prefixes"`
}

func (cf CloudFront__P) GET() ([]string, error) {
	const URL = "https://ip-ranges.amazonaws.com/ip-ranges.json"

	res, e := http.Get(URL)
	if e != nil {
		return nil, e
	}
	if res.StatusCode != http.StatusOK {
		return nil, e
	}

	cf.RAW, e = io.ReadAll(res.Body)
	if e != nil {
		return nil, e
	}

	e = json.Unmarshal(cf.RAW, &cf)
	if e != nil {
		return nil, e
	}
	cf.RAW = nil

	r := make([]string, 0)
	for _, ip := range cf.Prefixes {
		if len(ip.IP_pref) != 0 {
			r = append(r, ip.IP_pref)
		}
	}
	cf.Prefixes = nil

	return r, nil
}
