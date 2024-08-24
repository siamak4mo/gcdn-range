package providers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type CloudFront__P struct {
	RAW []byte
}

func (cf CloudFront__P) GET(cout ProvChan, flags int) error {
	defer func() {
		cf.RAW = nil
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

	switch flags {
	case DL_IPv4:
		d := struct {
			Prefexs []struct {
				IP_pref string `json:"ip_prefix"`
			} `json:"prefixes"`
		}{}

		e = json.Unmarshal(cf.RAW, &d)
		if e != nil {
			return e
		}

		for _, ip := range d.Prefexs {
			cout <- ip.IP_pref
		}
		break
	case DL_IPv6:
		d := struct {
			Prefexs_v6 []struct {
				IP_pref string `json:"ipv6_prefix"`
			} `json:"ipv6_prefixes"`
		}{}
		e = json.Unmarshal(cf.RAW, &d)
		if e != nil {
			return e
		}

		for _, ip := range d.Prefexs_v6 {
			cout <- ip.IP_pref
		}
		break
	case DL_ALL:
		d := struct {
			Prefexs []struct {
				IP_pref string `json:"ip_prefix"`
			} `json:"prefixes"`
			Prefexs_v6 []struct {
				IP_pref string `json:"ipv6_prefix"`
			} `json:"ipv6_prefixes"`
		}{}
		e = json.Unmarshal(cf.RAW, &d)
		if e != nil {
			return e
		}

		for _, ip := range d.Prefexs {
			cout <- ip.IP_pref
		}
		for _, ip := range d.Prefexs_v6 {
			cout <- ip.IP_pref
		}
		break
	}
	return nil
}
