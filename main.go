package main

import (
	"errors"
	"fmt"
	"gcdn_range/cmd"
	"os"
)

const (
	// output type
	STD_OUT = iota
	STD_FILE

	// output format
	FORMAT_RAW
	FORMAT_RAWV // verbose raw
	FORMAT_JSON
	FORMAT_CSV
	FORMAT_TSV

	// provider(s)
	PROV_ALL
	PROV_SELECTED
)

type config struct {
	output_t        uint8  // stdout or file
	output_filepath string // filepath

	format_o uint8 // raw,json,csv,tsv

	providers_t uint8    // all or selected
	providers   []string // provider(s) in selected mode
}

var cfg = config{
	output_t:    STD_OUT,
	format_o:    FORMAT_RAW,
	providers_t: PROV_ALL,
	providers:   make([]string, 0),
}

func arg_parser() error {
	argc := len(os.Args) - 1
	for i := 1; i < len(os.Args); i++ {
		argc -= 1
		switch os.Args[i] {
		case "--verbose", "--verbos", "--verb", "-v":
			if cfg.format_o == FORMAT_RAW {
				cfg.format_o = FORMAT_RAWV
			}

		case "--output", "--out", "-output", "-out", "-o":
			if argc >= 1 {
				i++
				argc--
				cfg.output_t = STD_FILE
				cfg.output_filepath = os.Args[i]
			} else {
				return errors.New("output flag needs an argument")
			}

		case "--format", "--form", "-f":
			if argc >= 1 {
				i++
				argc--
				if cfg.format_o == FORMAT_RAWV {
					continue
				}
				switch os.Args[i] {
				case "json", "j", "Json", "JSON", "J":
					cfg.format_o = FORMAT_JSON
				case "csv", "CSV", "Csv":
					cfg.format_o = FORMAT_CSV
				case "tsv", "TSV", "Tsv":
					cfg.format_o = FORMAT_TSV
				case "raw", "txt", "Raw", "RAW", "Txt", "TXT":
					cfg.format_o = FORMAT_RAW
				default:
					return errors.New(
						fmt.Sprintf("unknown format `%s`", os.Args[i]),
					)
				}
			} else {
				return errors.New("format flag needs an argument")
			}

		// --format shortcuts
		case "-json", "--json", "-j":
			cfg.format_o = FORMAT_JSON
		case "-csv", "--csv", "--cs", "-cs":
			cfg.format_o = FORMAT_CSV
		case "-tsc", "--tsv", "--ts", "-ts":
			cfg.format_o = FORMAT_TSV

		case "--providers", "--provider", "--prov", "-p":
			if argc >= 1 {
				i++
				argc--
				cfg.providers_t = PROV_SELECTED
				cfg.providers = append(cfg.providers, os.Args[i])
			} else {
				return errors.New("provider flag needs an argument")
			}

		case "--": // consider the rest as provider names
			if argc >= 1 {
				i++
				argc--
				cfg.providers_t = PROV_SELECTED
				cfg.providers = append(cfg.providers, os.Args[i:]...)
				return nil
			}

		default:
			return errors.New(
				fmt.Sprintf("unknown flag `%s` has been provided", os.Args[i]),
			)
		}
	}
	return nil
}

func not_imp(s string) {
	println(s + " not implemented yet!")
	os.Exit(1)
}

func main() {
	e := arg_parser()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v -- exiting.\n", e.Error())
		os.Exit(1)
	}

	dl := cmd.NewDownloader()
	if dl == nil {
		os.Exit(1)
	}

	switch cfg.format_o {
	case FORMAT_RAWV:
		dl = dl.SetOut(dl.Formated_RAW_Writer)
	case FORMAT_RAW:
		dl = dl.SetOut(dl.RAW_Writer)
	case FORMAT_CSV:
		not_imp("csv")
	case FORMAT_JSON:
		dl = dl.SetOut(dl.Json_Writer)
	case FORMAT_TSV:
		not_imp("tsv")
	}

	switch cfg.output_t {
	case STD_OUT:
		// defualt
	case STD_FILE:
		e := dl.SetOutPath(cfg.output_filepath)
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s -- Exiting.\n", e.Error())
			os.Exit(1)
		}
	}

	// TODO: add ip v4/v6 flags

	switch cfg.providers_t {
	case PROV_ALL:
		dl = dl.Init(cmd.DL_IPv4).DL_all()
	case PROV_SELECTED:
		dl = dl.Init(cmd.DL_IPv4).DL_names(cfg.providers)
	}

	dl.Do()

	defer dl.Done()
}
