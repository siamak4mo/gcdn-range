package providers

import (
	"encoding/json"
	"errors"
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

func (cf CloudFront__P) GET(cout ProvChan) error {
	defer func(){
		cf.RAW = nil
		cf.Prefixes = nil
		close(cout)
	}()
	const URL = "https://ip-ranges.amazonaws.com/ip-ranges.json"
	if cout == nil {
		return errors.New("Null Chanel")
	}

	res, e := http.Get(URL)
	if e != nil {
		return e
	}
	if res.StatusCode != http.StatusOK {
		return e
	}
	cf.RAW, e = io.ReadAll(res.Body)
	if e != nil {
		return e
	}
	e = json.Unmarshal(cf.RAW, &cf)
	if e != nil {
		return e
	}

	for _, ip := range cf.Prefixes {
		if len(ip.IP_pref) != 0 {
			cout <- ip.IP_pref
		}
	}
	return nil
}
