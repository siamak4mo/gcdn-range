package main

import (
	"gcdn_range/cmd"
)

func main() {
	d := cmd.NewDownloader()
	d = d.Init(cmd.DL_IPv6).DL_name("cloudflare").SetOut(cmd.Formatted_stdout).Do()
}
