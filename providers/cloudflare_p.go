package providers

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

type CloudFlare__P struct {
	RAW []byte
}

func (cf CloudFlare__P) GET(cout ProvChan) error {
	defer func() {
		cf.RAW = nil
		close(cout)
	}()
	const URL = "https://www.cloudflare.com/ips-v4"
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

	for _, v := range strings.Split(string(cf.RAW), "\n") {
		cout <- v
	}
	return nil
}
