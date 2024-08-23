package providers

import (
	"io"
	"net/http"
	"strings"
)

type CloudFlare__P struct {
	RAW []byte
}

func (cf CloudFlare__P) GET(result []string) ([]string, error) {
	const URL = "https://www.cloudflare.com/ips-v4"

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
	return strings.Split(string(cf.RAW), "\n"), nil
}
