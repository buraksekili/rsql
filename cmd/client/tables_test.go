package client

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestShowTables(t *testing.T) {
	db, mock := NewMockDB()
	defer db.Close()

	tests := []struct {
		name     string
		expected []string
	}{
		{
			name:     "show no table",
			expected: []string{},
		},
		{
			name:     "show single element table",
			expected: []string{"james"},
		},
		{
			name:     "show multiple element table",
			expected: []string{"king", "james"},
		},
	}

	// q is query which will be tested.
	q := `show tables`
	clientDB := getMockClientDB(db)

	for _, tc := range tests {
		var buff bytes.Buffer

		row := sqlmock.NewRows([]string{"name"})
		for _, table := range tc.expected {
			row.AddRow(table)
		}
		mock.ExpectQuery(q).WillReturnRows(row)

		tables := clientDB.showTables(&buff)
		if !reflect.DeepEqual(tc.expected, tables) {
			t.Fatalf("%s, got=%v expected=%v", tc.name, tables, tc.expected)
		}
	}
}
