package cmd

import (
	"fmt"
	"gcdn_range/providers"
	"os"
)

type Downloader struct {
	CIDR  []string
	Provs []providers.Provider
}

func NewDownloader() *Downloader {
	return &Downloader{}
}

func (dl *Downloader) Init() *Downloader {
	dl.CIDR = make([]string, 0)
	dl.Provs = nil

	return dl
}

func (dl *Downloader) do() {
	for _, p := range dl.Provs {
		if p.Pr != nil {
			tmp, e := p.Pr.GET(dl.CIDR)
			if e == nil {
				dl.CIDR = tmp

			} else {
				fmt.Fprintf(os.Stderr, "Could not download -- %v\n", e.Error())
			}
		}
	}
}

func (dl *Downloader) DL_all() *Downloader {
	dl.Provs = providers.CDNs
	dl.do()
	return dl
}

func (dl *Downloader) DL_name(name string) *Downloader {
	p, e := providers.SearchCDN(name)
	if e == nil {
		dl.Provs = []providers.Provider{p}
		dl.do()
	} else {
		fmt.Fprintf(os.Stderr, "%s -- %s", p.Name, e.Error())
	}
	return dl
}

func (dl *Downloader) DL_names(names []string) *Downloader {
	dl.Provs = providers.MkCDN(names)
	dl.do()
	return dl
}
