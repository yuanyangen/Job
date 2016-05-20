package mixed

import "testing"

func Test_Add(t *testing.T) {
	a := "12345"
	b := "123456"
	c := Add(a, b)
	if c != "135801" {
		t.Errorf("failed")
	}
}

func Test_Multiple(t *testing.T) {
	a := "12345"
	b := "1200000003456"
	c := Multiple(a, b)
	if c != "14814000042664320" {
		t.Errorf(" multiple failed")
	}
}
