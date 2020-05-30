package gorm

import (
	"testing"

	dbassert "github.com/hashicorp/dbassert"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func Test_FieldDomain(t *testing.T) {
	t.Parallel()
	cleanup, conn, _ := dbassert.TestSetup(t, "postgres")
	defer func() {
		if err := cleanup(); err != nil {
			t.Error(err)
		}
		if err := conn.Close(); err != nil {
			t.Error(err)
		}
	}()

	mockery := new(dbassert.MockTesting)
	dbassert := New(mockery, conn, "postgres")

	dbassert.FieldDomain(&TestModel{}, "PublicId", "dbasserts_public_id")
	mockery.AssertNoError(t)

	mockery.Reset()
	dbassert.FieldDomain(&TestModel{}, "nullable", "dbasserts_public_id")
	mockery.AssertError(t)
}

func Test_FieldNullable(t *testing.T) {
	t.Parallel()
	cleanup, conn, _ := dbassert.TestSetup(t, "postgres")
	defer func() {
		if err := cleanup(); err != nil {
			t.Error(err)
		}
		if err := conn.Close(); err != nil {
			t.Error(err)
		}
	}()

	mockery := new(dbassert.MockTesting)
	dbassert := New(mockery, conn, "postgres")

	dbassert.FieldNullable(&TestModel{}, "Nullable")
	mockery.AssertNoError(t)

	mockery.Reset()
	dbassert.FieldNullable(&TestModel{}, "PublicId")
	mockery.AssertError(t)
}

func Test_FieldIsNull(t *testing.T) {
	t.Parallel()
	cleanup, conn, _ := dbassert.TestSetup(t, "postgres")
	defer func() {
		if err := cleanup(); err != nil {
			t.Error(err)
		}
		if err := conn.Close(); err != nil {
			t.Error(err)
		}
	}()
	assert := assert.New(t)
	db, err := gorm.Open("postgres", conn)
	assert.NoError(err)

	v := 1
	m := CreateTestModel(t, db, nil, &v)

	mockery := new(dbassert.MockTesting)
	dbassert := New(mockery, conn, "postgres")

	dbassert.FieldIsNull(&m, "Nullable")
	mockery.AssertNoError(t)

	mockery.Reset()
	dbassert.FieldIsNull(&m, "typeint")
	mockery.AssertError(t)
}

func Test_FieldNotNull(t *testing.T) {
	t.Parallel()
	cleanup, conn, _ := dbassert.TestSetup(t, "postgres")
	defer func() {
		if err := cleanup(); err != nil {
			t.Error(err)
		}
		if err := conn.Close(); err != nil {
			t.Error(err)
		}
	}()
	assert := assert.New(t)
	db, err := gorm.Open("postgres", conn)
	assert.NoError(err)

	v := 1
	m := CreateTestModel(t, db, nil, &v)

	mockery := new(dbassert.MockTesting)
	dbassert := New(mockery, conn, "postgres")

	dbassert.FieldNotNull(&m, "Nullable")
	mockery.AssertError(t)

	mockery.Reset()
	dbassert.FieldNotNull(&m, "typeint")
	mockery.AssertNoError(t)
}