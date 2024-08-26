package providers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Incapsula__P struct {
	RAW []byte
}

func (cf Incapsula__P) GET(cout ProvChan, flags int) error {
	defer func() {
		cf.RAW = nil
		close(cout)
	}()
	const URL = "https://ip-ranges.amazonaws.com/ip-rangesdi.json"
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
		res.Body.Close()
		return e
	}

	switch flags {
	case DL_IPv4:
		d := struct {
			IPs []string `json:"ipRanges"`
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
			IPs []string `json:"ipv6Ranges"`
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
			IPs []string `json:"ipRanges"`
			IPv6 []string `json:"ipv6Ranges"`
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
