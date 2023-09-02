package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := GetHello("shiono")
	expect := "hello:shiono"
	if result != expect {
		t.Error("real:", result, "assert:", expect)
	}

	t.Log("TestHello")
}
