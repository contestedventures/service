package folder

import "testing"

func TestMethod(t *testing.T) {
	res := Method()
	if res != "test3" {
		t.Fail()
	}
}
