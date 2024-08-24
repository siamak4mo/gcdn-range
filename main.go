package main

import (
	"gcdn_range/cmd"
)

func main() {
	d := cmd.NewDownloader()

	d.Init().DL_name("cloudflare").Raw_stdout()
	d.Init().DL_name("test").Formatted_stdout()
}
