package providers

import (
	"errors"
	"net/http"
)

type Maxcdn__P struct {
	URL []string
}

func (cf Maxcdn__P) GET(cout ProvChan, flags int) error {
	defer func() {
		close(cout)
	}()

	switch flags {
	case DL_IPv4:
		cf.URL = []string{"https://support.maxcdn.com/hc/en-us/article_attachments/360051920551/maxcdn_ips.txt"}
	case DL_IPv6:
		cf.URL = []string{}
	case DL_ALL:
		cf.URL = []string{"https://support.maxcdn.com/hc/en-us/article_attachments/360051920551/maxcdn_ips.txt"}
	}
	if cout == nil {
		return errors.New("Null Chanel")
	}

	for _, url := range cf.URL {
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
