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

func (dl *Downloader) SetOutPath(path string) error {
	file, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		dl.Out.Writer = nil
		return err
	}
	dl.Out.Writer = file
	dl.Fout = file
	return nil
}

// generic writers
// each of these functions can be assigned to Downloader.Out.Writer
func (dl *Downloader) RAW_Writer(p *providers.Provider) {
	for cidr := range p.CIDR {
		fmt.Fprintln(dl.Out.Writer, cidr)
	}
}

func (dl *Downloader) Formated_RAW_Writer(p *providers.Provider) {
	if p.DLerr != nil {
		return
	}
	fmt.Fprintf(dl.Out.Writer, "%v:\n", p.Name)
	for cidr := range p.CIDR {
		fmt.Fprintf(dl.Out.Writer, "  - %s\n", cidr)
	}
}
