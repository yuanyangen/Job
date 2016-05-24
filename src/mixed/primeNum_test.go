package mixed

import "testing"

func Test_Prime(t *testing.T) {
	n := 100000000
	ret := Prime(n)
	if ret[0] != 1 || ret[1] != 2 || ret[2] != 3 {
		t.Errorf("get prime failed")
	}
}

