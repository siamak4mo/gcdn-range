package main

import (
	"fmt"
	"gcdn_range/cmd"
)

func main() {
	d := cmd.NewDownloader()

	d.Init().DL_all()
	fmt.Printf("%v\n", d.CIDR)
}
