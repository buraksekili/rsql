package client

import (
	"bytes"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDisplayDBStats(t *testing.T) {
	db, _ := NewMockDB()
	defer db.Close()

	client := getMockClientDB(db)

	var buff bytes.Buffer
	client.displayDBStats(&buff)
	got := strings.TrimSpace(buff.String())

	expOut := `===== STATS =====
Max Open Connections:	 0
pen Connection:		 1
Idle:			 1
In Use:			 0`

	if diff := cmp.Diff(expOut, got); diff != "" {
		t.Fatalf("mismatch DisplayDBStats output: diff %s", diff)
	}

	stats := getDBStatus(db)
	if stats.InUse != 0 {
		t.Fatalf("invalid InUse; got %d expected 0", stats.InUse)
	}

	if stats.Idle != 1 {
		t.Fatalf("invalid Idle; got %d expected 1", stats.Idle)
	}

	if stats.OpenConnections != 1 {
		t.Fatalf("invalid OpenConnections; got %d expected 1", stats.OpenConnections)
	}
}
