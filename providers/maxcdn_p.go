package providers

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

type Maxcdn__P struct {
	URL []string
	RAW []byte
}

func (cf Maxcdn__P) GET(cout ProvChan, flags int) error {
	defer func() {
		cf.RAW = nil
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
