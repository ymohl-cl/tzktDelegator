package pgsql

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestPGSQL_Driver(t *testing.T) {
	t.Run("should return the driver", func(t *testing.T) {
		db, _, err := sqlmock.New()
		if err != nil {
			t.Fatal(err)
		}
		pg := pgsql{
			driver: db,
		}

		driver := pg.Driver()
		if assert.NotNil(t, driver) {
			assert.Equal(t, driver, db)
		}
	})
}

func TestPGSQL_Close(t *testing.T) {
	t.Run("should close the driver", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal(err)
		}
		pg := pgsql{
			driver: db,
		}

		mock.ExpectClose()
		err = pg.Close()
		if assert.Nil(t, err) {
			assert.Nil(t, mock.ExpectationsWereMet())
		}
	})
	t.Run("should close the driver with an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal(err)
		}
		pg := pgsql{
			driver: db,
		}

		errExpect := errors.New("case error")
		mock.ExpectClose().WillReturnError(errExpect)
		err = pg.Close()
		if assert.Error(t, err) {
			assert.ErrorIs(t, err, errExpect)
			assert.Nil(t, mock.ExpectationsWereMet())
		}
	})
}
