package providers

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

type CloudFlare__P struct {
	URL []string
	RAW []byte
}

func (cf CloudFlare__P) GET(cout ProvChan, flags int) error {
	defer func() {
		cf.RAW = nil
		close(cout)
	}()

	switch flags {
	case DL_IPv4:
		cf.URL = []string{"https://www.cloudflare.com/ips-v4"}
	case DL_IPv6:
		cf.URL = []string{"https://www.cloudflare.com/ips-v6"}
	case DL_ALL:
		cf.URL = []string{"https://www.cloudflare.com/ips-v4",
			"https://www.cloudflare.com/ips-v6"}
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

		cf.RAW, e = io.ReadAll(res.Body)
		if e != nil {
			return e
		}

		for _, v := range strings.Split(string(cf.RAW), "\n") {
			cout <- v
		}
	}
	return nil
}
