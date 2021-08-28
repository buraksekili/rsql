package commands

import (
	"testing"
)

func TestNewDbClient(t *testing.T) {
	c := NewDbClient(nil)
	if c == nil {
		t.Fatalf("incorrent dbClient type; got %T expected %T", c, &DbClient{})
	}
}
