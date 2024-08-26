package providers

import (
	"errors"
	"net/http"
)

type Cachefly__P struct {
	URL []string
}

func (cf Cachefly__P) GET(cout ProvChan, flags int) error {
	defer func() {
		close(cout)
	}()

	switch flags {
	case DL_IPv4:
		cf.URL = []string{"https://cachefly.cachefly.net/ips/rproxy.txt"}
	case DL_IPv6:
		cf.URL = []string{}
	case DL_ALL:
		cf.URL = []string{"https://cachefly.cachefly.net/ips/rproxy.txt"}
	}
	if cout == nil {
		return errors.New("Null Chanel")
	}

	for _,url := range cf.URL {
		res, e := http.Get(url)
		if e != nil {
			return e
		}
		if res.StatusCode != http.StatusOK {
			return e
		}

		rw2cahn_from_http(res, cout)
	}
	return nil
}
