package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependence(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funfa caras")
	}
	t.Fail()
}
