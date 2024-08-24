package cmd

import (
	"fmt"
	"gcdn_range/providers"
	"os"
)

type Downloader struct {
	Provs []*providers.Provider
}

func NewDownloader() *Downloader {
	return &Downloader{}
}

func (dl *Downloader) Init() *Downloader {
	dl.Provs = nil
	return dl
}

func (dl *Downloader) do() {
	for _, p := range dl.Provs {
		if p.Pr != nil {
			p.DoFetch()
			if p.DLerr != nil {
				fmt.Fprintf(os.Stderr, "Could not download -- %v\n", p.DLerr)
			}
		}
	}
}

func (dl *Downloader) DL_all() *Downloader {
	//	dl.Provs = make()//
	dl.do()
	return dl
}

func (dl *Downloader) DL_name(name string) *Downloader {
	p, e := providers.SearchCDN(name)
	dl.Provs = []*providers.Provider{&p}

	if e == nil {
		dl.do()
	} else {
		dl.Provs[0].DLerr = e
		fmt.Fprintf(os.Stderr, "%s -- %v\n", name, e.Error())
	}
	return dl
}

func (dl *Downloader) DL_names(names []string) *Downloader {
	dl.Provs = providers.MkCDN(names)
	dl.do()
	return dl
}
