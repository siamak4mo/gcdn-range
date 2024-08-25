package cmd

import (
	"fmt"
	"gcdn_range/providers"
	"os"
)

func (dl *Downloader) SetOut(dw Dl_Writer) *Downloader {
	dl.Out = OutputWriter{
		Do:     dw,
		Writer: os.Stdout,
	}
	return dl
}

func (*Downloader) Raw_stdout(p *providers.Provider) {
// generic writers
// each of these functions can be assigned to Downloader.Out.Writer
func (dl *Downloader) Raw_stdout(p *providers.Provider) {
	for cidr := range p.CIDR {
		fmt.Fprintln(dl.Out.Writer, cidr)
	}
}

func (dl *Downloader) Formatted_stdout(p *providers.Provider) {
	if p.DLerr != nil {
		return
	}
	fmt.Fprintf(dl.Out.Writer, "%v:\n", p.Name)
	for cidr := range p.CIDR {
		fmt.Fprintf(dl.Out.Writer, "  - %s\n", cidr)
	}
}

// TODO:  file output
//        json stdout
