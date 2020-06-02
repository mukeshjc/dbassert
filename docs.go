// Package dbassert provides a set of assertions for testing Go database
// applications.
//
// Example Usage:
//
// import (
// 	"testing"
//
// 	"github.com/hashicorp/dbassert"
// )
//
// func TestSomeDatabase(t *testing.T) {
//	db, err := sql.Open("postgres", "postgres://postgres:secret@localhost:%s?sslmode=disable")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer db.Close()
// 	dbassert := dbassert.New(t, conn, "postgres")
// 	dbassert.Nullable("some_table_name", "some_column")
// }

package dbassert
