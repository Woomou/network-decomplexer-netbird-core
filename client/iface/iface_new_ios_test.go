//go:build ios

package iface

import (
	"os"
	"testing"

	"github.com/netbirdio/netbird/client/iface/netstack"
)

func TestIOSNetstackModeIsExplicitlyEnabled(t *testing.T) {
	old := os.Getenv(netstack.EnvUseNetstackMode)
	t.Cleanup(func() { _ = os.Setenv(netstack.EnvUseNetstackMode, old) })
	if err := os.Setenv(netstack.EnvUseNetstackMode, "true"); err != nil {
		t.Fatal(err)
	}
	if !netstack.IsEnabled() {
		t.Fatal("expected iOS Netstack mode to be enabled")
	}
}
