package main

import (
	"os"
	"testing"

	"github.com/usace/wat-go"
	"gopkg.in/yaml.v3"
)

func TestReadPayload(t *testing.T) {
	path := "./exampledata/payload.json"
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fail()
	}
	mp := wat.Payload{}
	err = yaml.Unmarshal(b, &mp)
	if err != nil {
		t.Fail()
	}
}
func TestComputePayload(t *testing.T) {
	path := "./exampledata/payload.json"
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fail()
	}
	mp := wat.Payload{}
	err = yaml.Unmarshal(b, &mp)
	if err != nil {
		t.Fail()
	}
	err = computePayload(mp)
	if err != nil {
		t.Fail()
	}
}
