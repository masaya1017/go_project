package test
import (
    "testing"
)

func capital(expected string) string {
    if expected == "Hello" {
        return expected
    } else {
        return expected + "1"
    }
}

func TestHelloOK(t *testing.T) {
    //想定内の文字であるpattern
    expected := "Hello"
    if got := capital(expected); got != expected {
        t.Errorf("got %q, want %q", got, expected)
    }
}

func TestHelloNG(t *testing.T) {
    //失敗pattern
    expected := "Hello1"
    got := capital(expected)
    if got != expected {
        t.Errorf("expected: %q, got: %q", expected, got)
    }
}
