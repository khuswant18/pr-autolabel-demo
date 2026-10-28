package scan

import "testing"

func TestRun(t *testing.T) {
	if got, want := Run("cluster"), "scanned cluster: 0 failed controls"; got != want {
		t.Fatalf("Run() = %q, want %q", got, want)
	}
}
