package cmd

import (
	"fmt"
	"gcdn_range/providers"
	"os"
)

type Dl_Writer func(*providers.Provider)

type Downloader struct {
	Provs   []*providers.Provider
	ipFlags int
	Out     OutputWriter
}

type OutputWriter struct {
	Do   Dl_Writer
	Path *string
}

const (
	DL_IPv4 = iota
	DL_IPv6
	DL_ALL
)

func NewDownloader() *Downloader {
	return &Downloader{}
}

func (dl *Downloader) Init(flags int) *Downloader {
	dl.Provs = nil
	dl.ipFlags = flags
	return dl
}

func (dl *Downloader) Do() *Downloader {
	for _, p := range dl.Provs {
		p.CIDR = make(providers.ProvChan)
		go p.DoFetch(dl.ipFlags)
		dl.Out.Do(p)
		if p.DLerr != nil {
			fmt.Fprintf(os.Stderr, "Could not download %s -- %v\n", p.Name, p.DLerr)
		}
	}
	return dl
}

func (dl *Downloader) DL_all() *Downloader {
	dl.Provs = providers.MKProvs()
	return dl
}

func (dl *Downloader) DL_name(name string) *Downloader {
	p, e := providers.SearchCDN(name)
	dl.Provs = []*providers.Provider{&p}

	if e != nil {
		dl.Provs[0].DLerr = e
		fmt.Fprintf(os.Stderr, "%s -- %v\n", name, e.Error())
	}
	return dl
}

func (dl *Downloader) DL_names(names []string) *Downloader {
	dl.Provs = providers.MkProv(names)
	return dl
}
