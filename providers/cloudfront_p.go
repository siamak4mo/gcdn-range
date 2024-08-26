package providers

import (
	"encoding/json"
	"errors"
	"net/http"
)

type CloudFront__P struct {
	RAW []byte
}

type IP_Pref struct {
	IP string `json:"ip_prefix"`
}
type IP_Pref6 struct {
	IP string `json:"ipv6_prefix"`
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
		res.Body.Close()
		return e
	}
	dec := json.NewDecoder(res.Body)

	switch flags {
	case DL_IPv4:
		nextArrayToken(dec) // go to the first `[` beginning of ipv4
		for dec.More() {
			var data IP_Pref
			e = dec.Decode(&data)
			if e != nil {
				break
			}
			cout <- data.IP
		}

	case DL_IPv6:
		nextArrayToken(dec) // beginning of ipv4
		nextArrayToken(dec) // beginning of ipv6
		var data6 IP_Pref6
		for dec.More() {
			e = dec.Decode(&data6)
			if e != nil {
				break
			}
			cout <- data6.IP
		}

	case DL_ALL:
		nextArrayToken(dec) // beginning of ipv4
		var (
			data  IP_Pref
			data6 IP_Pref6
		)
		for dec.More() {
			e = dec.Decode(&data)
			if e != nil {
				break
			}
			cout <- data.IP
		}
		nextArrayToken(dec) // beginning of ipv6
		for dec.More() {
			e = dec.Decode(&data6)
			if e != nil {
				break
			}
			cout <- data6.IP
		}
	}
	return nil
}
