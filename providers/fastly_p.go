package providers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Fastly__P struct {
	RAW []byte
}

func (cf Fastly__P) GET(cout ProvChan, flags int) error {
	defer func() {
		cf.RAW = nil
		close(cout)
	}()
	const URL = "https://api.fastly.com/public-ip-list"
	if cout == nil {
		return errors.New("Null Chanel")
	}

	res, e := http.Get(URL)
	if e != nil {
		return e
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return e
	}
	cf.RAW, e = io.ReadAll(res.Body)
	if e != nil {
		res.Body.Close()
		return e
	}

	switch flags {
	case DL_IPv4:
		d := struct {
			IPs []string `json:"addresses"`
		}{}

		e = json.Unmarshal(cf.RAW, &d)
		if e != nil {
			return e
		}

		for _, ip := range d.IPs {
			cout <- ip
		}
		break
	case DL_IPv6:
		d := struct {
			IPs []string `json:"ipv6_addresses"`
		}{}
		e = json.Unmarshal(cf.RAW, &d)
		if e != nil {
			return e
		}

		for _, ip := range d.IPs {
			cout <- ip
		}
		break
	case DL_ALL:
		d := struct {
			IPs []string `json:"addresses"`
			IPv6 []string `json:"ipv6_addresses"`
		}{}
		e = json.Unmarshal(cf.RAW, &d)
		if e != nil {
			return e
		}

		for _, ip := range d.IPs {
			cout <- ip
		}
		for _, ip := range d.IPv6 {
			cout <- ip
		}
		break
	}
	return nil
}
