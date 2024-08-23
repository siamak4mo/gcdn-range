package providers

import (
	"testing"
)

func (p Provider) Equal(q Provider) bool {
	if p.ID == q.ID &&
		p.Name == q.Name {
		return true
	}
	return false
}

func Test__CDNs_integrity(t *testing.T) {
	for idx, c := range CDNs {
		if int(c.ID) != idx {
			t.Errorf("%s\n%s\nCDNs[%v]->ID  !=  %v\n",
				"CDNs data integrity test falied",
				"make sure the ID of Providers is equal to their index in the CDNs array",
				c.ID, idx)
		}
	}
}
