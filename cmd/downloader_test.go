package cmd

import (
	"gcdn_range/providers"
	"testing"
)

const (
	WRONG_PROV_LIST = "Wrong provider list"
)

func Test_Dl_all(t *testing.T) {
	d := NewDownloader()
	d.DL_all()

	if len(d.Provs) != len(providers.CDNs) {
		t.Errorf("Dl_all failed -- %s\n", WRONG_PROV_LIST)
	}
	for idx, p := range d.Provs {
		if p.Name != providers.CDNs[idx].Name {
			t.Errorf("Dl_all failed -- %s\n", WRONG_PROV_LIST)
		}
	}
}
