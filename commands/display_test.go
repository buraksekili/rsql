package commands

import (
	"bytes"
	"database/sql"
	"log"
	"os"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/buraksekili/rsql/data"
	"github.com/buraksekili/selog"
	_ "github.com/go-sql-driver/mysql"
)

func TestDisplayTable(t *testing.T) {
	db, mock := NewMockDB()
	defer db.Close()

	tests := []struct {
		name     string
		expected [][]string
	}{
		{
			name:     "display empty table",
			expected: [][]string{},
		},
		{
			name:     "display no element table",
			expected: [][]string{{""}},
		},
		{
			name:     "display single empty element table",
			expected: [][]string{{" "}},
		},
		{
			name:     "display single element table",
			expected: [][]string{{"Burak"}},
		},
		{
			name:     "display multiple element table",
			expected: [][]string{{"Burak"}, {"rsql"}},
		},
	}

	// q is query which will be tested.
	q := `select * from test;`
	clientDB := getMockClientDB(db)

	for _, tc := range tests {
		var buff bytes.Buffer

		rows := sqlmock.NewRows([]string{"name"})
		for _, r := range tc.expected {
			for _, rr := range r {
				rows.AddRow(rr)
			}
		}
		mock.ExpectQuery(regexp.QuoteMeta(q)).WillReturnRows(rows)

		content, err := clientDB.displayTable("test", &buff)
		if err != nil {
			t.Fatalf("error while displaying table: %s", err)
		}
		if !reflect.DeepEqual(tc.expected, content) {
			t.Fatalf("%s, got=%v expected=%v", tc.name, content, tc.expected)
		}
	}

}

func NewMockDB() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func getMockClientDB(db *sql.DB) *DbClient {
	logger := log.New(os.Stdout, "rsql ", log.LstdFlags|log.Lshortfile)
	l := selog.NewLogger(logger)
	return &DbClient{&data.ConnInfo{}, l, db}
}
