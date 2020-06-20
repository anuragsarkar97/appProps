package main

import (
	"AppProps/src"
	"testing"
)

func TestIntegration(t *testing.T) {
	config1 := src.UseProps("resources/application.properties") // should pass
	src.UseProps("resources/application1.properties")           // should fail
	src.UseProps("resources/application1.not-properties")       // should fail
	if config1.Get("config-1") != "response1" {
		t.Errorf("expected answer is %s, answer loaded %s", "response1", config1.Get("config-1"))
	}
	if config1.Get("config-2") != "" {
		t.Errorf("expected answer is %s, answer loaded %s", "", config1.Get("config-1"))
	}
}