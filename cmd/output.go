package cmd

import (
	"fmt"
	"gcdn_range/providers"
)

func (dl *Downloader) SetOut(dw Dl_Writer) *Downloader {
	dl.Out = OutputWriter{
		Do: dw,
	}
	return dl
}

func (*Downloader) Raw_stdout(p *providers.Provider) {
	for cidr := range p.CIDR {
		fmt.Println(cidr)
	}
}

func (*Downloader) Formatted_stdout(p *providers.Provider) {
	if p.DLerr != nil {
		return
	}
	fmt.Printf("%v:\n", p.Name)
	for cidr := range p.CIDR {
		fmt.Printf("  - %s\n", cidr)
	}
}

// TODO:  file output
//        json stdout
