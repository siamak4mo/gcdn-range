package cmd

import "fmt"

func (dl *Downloader) Raw_stdout() *Downloader {
	for _,p := range dl.Provs {
		for _,cidr := range p.CIDR {
			fmt.Println(cidr)
		}
	}
	return dl
}

func (dl *Downloader) Formatted_stdout() *Downloader {
	for _,p := range dl.Provs {
		if p.DLerr != nil {
			continue
		}
		fmt.Printf("%v:\n", p.Name)
		if len(p.CIDR) == 0 {
			fmt.Println("  - Empty")
			continue
		}
		for _,cidr := range p.CIDR {
			fmt.Println(cidr)
		}
	}
	return dl
}

// TODO:  file output
//        json stdout
