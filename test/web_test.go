package test

import (
	"fmt"
	"go_project/lib"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ここで必要に応じてパスを指定します
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")

}

func TestHelloHandler(t *testing.T) {
	// モックのテスト
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}
	lib.SayHello(t)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloHandler)

	handler.ServeHTTP(rr, req)

	t.Log("Response body:", rr.Body.String())

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := "Hello, World!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
