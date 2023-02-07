package internal

import (
	"gee/network"
	"testing"
)

func TestServer(t *testing.T) {
	s := network.NewServer("Gee v0.0.1")
	s.Run()
}
