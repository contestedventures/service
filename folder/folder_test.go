package folder

import "testing"

func TestMethod(t *testing.T) {
	res := Method()
	if res != "test2" {
		t.Fail()
	}
}
