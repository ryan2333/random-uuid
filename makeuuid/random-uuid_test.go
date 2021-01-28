package makeuuid

import "testing"

func TestRandomUuid(t *testing.T) {
	r := RandomUuid(20)
	if len(r) != 20 {
		t.Errorf("get uuid failed, uuid len: %d", len(r))
		return
	}
	t.Log("uuid: ", r)
}
