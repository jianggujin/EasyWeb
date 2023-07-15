package server

import (
	"os"
	"testing"
)

func TestStartServer(t *testing.T) {
	dir, _ := os.Getwd()
	t.Log(dir)
	//StartServer()
}
