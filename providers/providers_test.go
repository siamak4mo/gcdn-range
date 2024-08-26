package providers

import (
	"testing"
)

func (p Provider) Equal(q Provider) bool {
	if p.id == q.id &&
		p.Name == q.Name {
		return true
	}
	return false
}

func Test__CDNs_integrity(t *testing.T) {
	for idx, c := range CDNs {
		if int(c.id) != idx {
			t.Errorf("%s\n%s\nCDNs[%v]->ID  !=  %v\n",
				"CDNs data integrity test falied",
				"make sure the ID of Providers is equal to their index in the CDNs array",
				c.id, idx)
		}
	}
}

func Test_MKProvs(t *testing.T) {
	idx := []int{0, 2, len(CDNs) - 1}
	names := make([]string, len(idx))

	for i := range idx {
		names[i] = CDNs[i].Name
	}
	p := MkProv(names)

	if len(idx) != len(p) {
		t.Errorf("MKProv test failed, wrong provider list\n")
	}
	for i := range idx {
		if names[i] != p[i].Name {
			t.Errorf("MKProv test failed, wrong provider list\n")
		}
	}
}

func Test_SearchCDN(t *testing.T) {
	for i,c := range CDNs {
		s_c, e := SearchCDN(c.Name)
		if s_c.Name != c.Name ||
			int(s_c.id) != i || e != nil {
			t.Errorf("SearchCDN failed, wrong name or ID\n")
		}
	}

	if _, e := SearchCDN("!!!"); e == nil {
		t.Errorf("SearchCDN nil test failed")
	}
}
