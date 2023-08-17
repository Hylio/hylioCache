package hyliocache

import "testing"

func TestCheckAddr(t *testing.T) {
	a := "localhost:8080"
	b := "192.168.0.1:443"
	c := "259.168.0.2:sss"
	if resA := CheckAddr(a); !resA {
		t.Errorf("A is wrong")
	}
	if resB := CheckAddr(b); !resB {
		t.Errorf("B is wrong")
	}
	if resC := CheckAddr(c); resC {
		t.Errorf("C is wrong")
	}
}
