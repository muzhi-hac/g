package gee

import "testing"

func TestNestedGroup(t *testing.T) {
	r := New()

	v1 := r.Group("/v1")
	v2 := v1.Group("/v2")
	v3 := v2.Group("/v3")
	//v1.GET("")
	v1.addRoute("GET", "/hello/:name", nil)
	//v1.GET("GET", "/hello/:name")

	if v2.prefix != "/v1/v2" {
		t.Fatal("v2 prefix should be /v1/v2")
	}
	if v3.prefix != "/v1/v2/v3" {
		t.Fatal("v2 prefix should be /v1/v2")
	}
}
